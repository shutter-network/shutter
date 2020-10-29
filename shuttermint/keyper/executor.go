package keyper

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/brainbot-com/shutter/shuttermint/contract"
)

const (
	skipCheckInterval    = 5 * time.Second
	halfStepPollInterval = 1 * time.Second
	kickOffBlockStagger  = 5
)

// CipherHalfStep returns the half step number of the cipher execution step.
func (p *CipherExecutionParams) CipherHalfStep() uint64 {
	return 2 * p.BatchIndex
}

// PlainHalfStep returns the half step number of the plain execution step.
func (p *CipherExecutionParams) PlainHalfStep() uint64 {
	return 2*p.BatchIndex + 1
}

// Run runs the executor.
func (ex *Executor) Run() error {
	for {
		select {
		case p := <-ex.cipherExecutionParams:
			err := ex.fastForward(p.BatchIndex, true)
			if err != nil {
				log.Printf("Error fast-forwarding to batch #%d: %s", p.BatchIndex, err)
				continue
			}

			err = ex.executeBatch(p)
			if err != nil {
				log.Printf("Error executing batch #%d: %s", p.BatchIndex, err)
				continue
			}

		case <-time.After(skipCheckInterval):
			err := ex.fastForward(math.MaxUint64, false)
			if err != nil {
				log.Printf("Error fast-forwarding: %s", err)
				continue
			}

		case <-ex.ctx.Done():
			return nil
		}
	}
}

// executeBatch executes the current batch (both cipher and plain parts) using the given
// parameters.
func (ex *Executor) executeBatch(p CipherExecutionParams) error {
	batchParams, err := ex.cc.ConfigContract.QueryBatchParams(ex.callOpts(nil), p.BatchIndex)
	if err != nil {
		return fmt.Errorf("error querying batch params for batch %d: %w", p.BatchIndex, err)
	}

	kickOffBlock, err := ex.cipherExecutionKickOffBlock(batchParams)
	if err != nil {
		return err
	}
	blockReached, err := ex.waitForBlockOrHalfStep(kickOffBlock, p.CipherHalfStep()+1)
	if err != nil {
		return err
	}
	if !blockReached {
		// block was not reached, so half step was reached, so someone else has executed the batch
		return nil
	}

	err = ex.executeCipherHalfStep(p)
	if err != nil {
		return err
	}

	err = ex.executePlainHalfStep(batchParams)
	if err != nil {
		return err
	}

	return nil
}

// fastForward skips all cipher batches that are skippable right now and executes all plain
// batches. It will stop when the last batch before untilBatchIndex has been handled. If wait is
// false, it will return as soon as immediate progress cannot be made. Otherwise, it will wait.
func (ex *Executor) fastForward(untilBatchIndex uint64, wait bool) error {
	for {
		numHalfSteps, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(ex.callOpts(nil))
		if err != nil {
			return fmt.Errorf("error querying current number of half steps: %w", err)
		}
		isCipher := numHalfSteps%2 == 0
		batchIndex := numHalfSteps / 2

		if batchIndex >= untilBatchIndex {
			return nil
		}

		batchParams, err := ex.cc.ConfigContract.QueryBatchParams(ex.callOpts(nil), batchIndex)
		if err != nil {
			return fmt.Errorf("error querying batch params for batch %d: %w", batchIndex, err)
		}

		var kickOffBlock uint64
		if isCipher {
			kickOffBlock, err = ex.cipherSkipKickOffBlock(batchParams)
		} else {
			kickOffBlock, err = ex.plainExecutionKickOffBlock(batchParams)
		}
		if err != nil {
			return err
		}

		var blockReached bool
		if !wait {
			blockReached, err = ex.isBlockReached(kickOffBlock)
		} else {
			blockReached, err = ex.waitForBlockOrHalfStep(kickOffBlock, numHalfSteps+1)
		}
		if err != nil {
			return err
		}
		if !blockReached {
			return nil
		}

		if isCipher {
			err = ex.skipCipherHalfStep(batchParams)
		} else {
			err = ex.executePlainHalfStep(batchParams)
		}
		if err != nil {
			return err
		}
	}
}

// executePlainHalfStep executes the plain portion of the current batch and returns when it's
// done. It does not perform any prior checks.
func (ex *Executor) executePlainHalfStep(batchParams BatchParams) error {
	batch, err := ex.fetchPlainBatch(batchParams)
	if err != nil {
		return err
	}

	tx, err := ex.cc.ExecutorContract.ExecutePlainBatch(ex.transactOpts(), batch)
	if err != nil {
		return fmt.Errorf("error sending plain batch execution tx: %w", err)
	}
	log.Printf("Plain batch execution tx sent for batch #%d with %d txs: %s", batchParams.BatchIndex, len(batch), tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, (batchParams.BatchIndex+1)*2)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %w", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to execute plain half step", tx.Hash().Hex())
		}
		log.Printf("Plain batch execution tx for batch #%d successful", batchParams.BatchIndex)
	} else {
		log.Printf("Plain batch execution tx for batch #%d potentially redundant", batchParams.BatchIndex)
	}

	return nil
}

// fetchPlainBatch downloads the plain txs for the given batch.
// This method should only be called once the batching period is over, otherwise some
// transactions might be missed.
func (ex *Executor) fetchPlainBatch(batchParams BatchParams) ([][]byte, error) {
	opts := ex.filterOpts(batchParams.StartBlock, &batchParams.EndBlock)
	it, err := ex.cc.BatcherContract.FilterTransactionAdded(opts)
	if err != nil {
		return [][]byte{}, fmt.Errorf("error registering TransactionAdded filter to download plain batch: %w", err)
	}

	batch := [][]byte{}
	for it.Next() {
		event := it.Event
		if event.TransactionType == contract.TransactionTypePlain {
			batch = append(batch, event.Transaction)
		}
	}
	it.Close()

	if it.Error() != nil {
		return [][]byte{}, fmt.Errorf("error fetching plain batch: %w", it.Error())
	}

	return batch, nil
}

// skipCipherHalfStep skips the current cipher half step and returns when it's done. It does not
// perform any prior checks.
func (ex *Executor) skipCipherHalfStep(batchParams BatchParams) error {
	tx, err := ex.cc.ExecutorContract.SkipCipherExecution(ex.transactOpts())
	if err != nil {
		return fmt.Errorf("error sending skip cipher execution tx: %w", err)
	}
	log.Printf("Cipher execution skip tx sent for batch #%d: %s", batchParams.BatchIndex, tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, batchParams.BatchIndex*2+1)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %w", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to skip cipher half step", tx.Hash().Hex())
		}
		log.Printf("Cipher half step of batch #%d successfully skipped in tx %s", batchParams.BatchIndex, tx.Hash().Hex())
	} else {
		log.Printf("Cipher half step of batch #%d skipped by another keyper", batchParams.BatchIndex)
	}

	return nil
}

// executeCipherHalfStep executes the current cipher half step and returns when it's done. It does
// not perform any prior checks.
func (ex *Executor) executeCipherHalfStep(p CipherExecutionParams) error {
	tx, err := ex.cc.ExecutorContract.ExecuteCipherBatch2(
		ex.transactOpts(),
		p.CipherBatchHash,
		p.DecryptedTxs,
		p.DecryptionKey,
		p.DecryptionSignerIndices,
		p.DecryptionSignatures,
	)
	if err != nil {
		return fmt.Errorf("error sending cipher execution tx: %s", err)
	}
	log.Printf("Cipher batch execution tx sent for batch #%d with %d txs: %s", p.BatchIndex, len(p.DecryptedTxs), tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, p.CipherHalfStep()+1)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %s", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to execute cipher half step", tx.Hash().Hex())
		}
		log.Printf("Cipher half step of batch #%d successfully executed in tx %s", p.BatchIndex, tx.Hash().Hex())
	} else {
		log.Printf("Cipher half step of batch #%d executed by another keyper", p.BatchIndex)
	}

	return nil
}

// kickOffDelay returns the number of blocks to wait before sending a tx.
func (ex *Executor) kickOffDelay(batchParams BatchParams) (uint64, error) {
	keyperIndex, ok := batchParams.BatchConfig.KeyperIndex(ex.cc.Address())
	if !ok {
		return 0, fmt.Errorf("not a keyper %s", ex.cc.Address().Hex())
	}
	place := (batchParams.BatchIndex + keyperIndex) % uint64(len(batchParams.BatchConfig.Keypers))
	return place * kickOffBlockStagger, nil
}

// cipherExecutionKickOffBlock returns the block number at which the cipher half step should be
// executed.
func (ex *Executor) cipherExecutionKickOffBlock(batchParams BatchParams) (uint64, error) {
	kickOffDelay, err := ex.kickOffDelay(batchParams)
	if err != nil {
		return 0, err
	}

	return batchParams.EndBlock + kickOffDelay, nil
}

// plainExecutionKickOffBlock returns the block number at which the plain half step should be
// executed.
func (ex *Executor) plainExecutionKickOffBlock(batchParams BatchParams) (uint64, error) {
	cipherKickOff, err := ex.cipherExecutionKickOffBlock(batchParams)
	if err != nil {
		return 0, err
	}
	return cipherKickOff + 1, nil
}

// cipherSkipKickOffBlock returns the block number at which the cipher half step should be
// skipped.
func (ex *Executor) cipherSkipKickOffBlock(batchParams BatchParams) (uint64, error) {
	kickOffDelay, err := ex.kickOffDelay(batchParams)
	if err != nil {
		return 0, err
	}

	return batchParams.ExecutionTimeoutBlock + kickOffDelay, nil
}

func (ex *Executor) callOpts(blockNumber *big.Int) *bind.CallOpts {
	return &bind.CallOpts{
		Context:     ex.ctx,
		BlockNumber: blockNumber,
	}
}

func (ex *Executor) transactOpts() *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(ex.cc.signingKey)
	opts.Context = ex.ctx

	return opts
}

func (ex *Executor) watchOpts(blockNumber *big.Int) *bind.WatchOpts {
	n := blockNumber.Uint64()
	return &bind.WatchOpts{
		Context: ex.ctx,
		Start:   &n,
	}
}

func (ex *Executor) filterOpts(start uint64, end *uint64) *bind.FilterOpts {
	return &bind.FilterOpts{
		Context: ex.ctx,
		Start:   start,
		End:     end,
	}
}

// headerChannel creates a channel to which the current and all future block headers will be sent.
func (ex *Executor) headerChannel(ctx context.Context) (<-chan *types.Header, <-chan error) {
	headers := make(chan *types.Header)
	errors := make(chan error)

	go func() {
		defer close(headers)
		defer close(errors)

		headersIn := make(chan *types.Header)
		sub, err := ex.client.SubscribeNewHead(ctx, headersIn)
		if err != nil {
			errors <- fmt.Errorf("error subscribing to new heads: %w", err)
			return
		}
		defer sub.Unsubscribe()

		firstHeader, err := ex.client.HeaderByNumber(ctx, nil)
		if err != nil {
			errors <- fmt.Errorf("error querying current head: %w", err)
			return
		}
		headers <- firstHeader

		for {
			select {
			case <-ctx.Done():
				errors <- ctx.Err()
				return
			case err := <-sub.Err():
				errors <- fmt.Errorf("error while listening to new heads: %w", err)
				return
			case header := <-headersIn:
				if header.Hash() == firstHeader.Hash() {
					continue
				}
				headers <- header
			}
		}
	}()

	return headers, errors
}

// receiptChannel creates a channel to which the receipt for the given tx will be sent when it
// arrives.
func (ex *Executor) receiptChannel(ctx context.Context, tx *types.Transaction) (<-chan *types.Receipt, <-chan error) {
	receipts := make(chan *types.Receipt)
	errors := make(chan error)

	go func() {
		defer close(receipts)
		defer close(errors)

		receipt, err := bind.WaitMined(ctx, ex.client, tx)
		if err != nil {
			errors <- fmt.Errorf("error while waiting for receipt of tx %s: %w", tx.Hash().Hex(), err)
		} else {
			receipts <- receipt
		}
	}()

	return receipts, errors
}

// halfStepChannel creates a channel to which the current and all future half step numbers are
// sent.
func (ex *Executor) halfStepChannel(ctx context.Context) (<-chan uint64, <-chan error) {
	halfSteps := make(chan uint64)
	errors := make(chan error)

	go func() {
		defer close(halfSteps)
		defer close(errors)

		// send current half step
		header, err := ex.client.HeaderByNumber(ctx, nil)
		if err != nil {
			errors <- fmt.Errorf("error querying current head: %w", err)
			return
		}
		callOpts := &bind.CallOpts{
			Context:     ctx,
			BlockNumber: header.Number,
		}
		startHalfStep, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(callOpts)
		halfSteps <- startHalfStep

		// subscribe to all events that change half step
		watchOpts := ex.watchOpts(header.Number)

		batchExecutedEvents := make(chan *contract.ExecutorContractBatchExecuted)
		batchExecutedSub, err := ex.cc.ExecutorContract.WatchBatchExecuted(watchOpts, batchExecutedEvents)
		if err != nil {
			errors <- fmt.Errorf("error subscribing to BatchExecuted events: %w", err)
			return
		}
		defer batchExecutedSub.Unsubscribe()

		cipherSkippedEvents := make(chan *contract.ExecutorContractCipherExecutionSkipped)
		cipherSkippedSub, err := ex.cc.ExecutorContract.WatchCipherExecutionSkipped(watchOpts, cipherSkippedEvents)
		if err != nil {
			errors <- fmt.Errorf("error subscribing to CipherExecutionSkipped events: %w", err)
			return
		}
		defer cipherSkippedSub.Unsubscribe()

		// send new half step on each event
		for {
			select {
			case <-ctx.Done():
				errors <- ctx.Err()
				return
			case err := <-batchExecutedSub.Err():
				errors <- fmt.Errorf("error while listening for BatchExecuted events: %w", err)
				return
			case err := <-cipherSkippedSub.Err():
				errors <- fmt.Errorf("error while listening for CipherExecutionSkipped events: %w", err)
				return
			case event := <-batchExecutedEvents:
				halfSteps <- event.NumExecutionHalfSteps
			case event := <-cipherSkippedEvents:
				halfSteps <- event.NumExecutionHalfSteps
			}
		}
	}()

	return halfSteps, errors
}

// waitForReceiptOrHalfStep waits until the given tx is included in a block or until the given
// half step is executed, whatever happens first. The tx's receipt (in the former case) or nil (in
// the latter) is returned.
func (ex *Executor) waitForReceiptOrHalfStep(tx *types.Transaction, n uint64) (*types.Receipt, error) {
	ctx, cancel := context.WithCancel(ex.ctx)
	defer cancel()

	receipts, receiptErrors := ex.receiptChannel(ctx, tx)
	halfSteps, halfStepErrors := ex.halfStepChannel(ctx)
	for {
		select {
		case <-ex.ctx.Done():
			return nil, ex.ctx.Err()
		case err := <-receiptErrors:
			return nil, err
		case err := <-halfStepErrors:
			return nil, err
		case r := <-receipts:
			return r, nil
		case halfStep := <-halfSteps:
			if halfStep < n {
				continue
			}

			// receipt and half step might arrive at the same time, but we prefer to return
			// the receipt. So wait briefly to check if a receipt is available.
			select {
			case r := <-receipts:
				return r, nil
			case <-time.After(500 * time.Millisecond):
				return nil, nil
			case <-ex.ctx.Done():
				return nil, nil
			}
		}
	}
}

// waitForBlockOrHalfStep waits for the given block number or the given half step, whichever comes
// first. It returns true if the block is reached first and false if the half step is reached
// first.
func (ex *Executor) waitForBlockOrHalfStep(blockNumber uint64, halfStep uint64) (bool, error) {
	ctx, cancel := context.WithCancel(ex.ctx)
	defer cancel()

	headers, headerErrors := ex.headerChannel(ctx)
	halfSteps, halfStepErrors := ex.halfStepChannel(ctx)

	for {
		select {
		case h := <-headers:
			if h.Number.Uint64() >= blockNumber {
				return true, nil
			}
		case s := <-halfSteps:
			if s >= halfStep {
				return false, nil
			}
		case err := <-halfStepErrors:
			return false, err
		case err := <-headerErrors:
			return false, err
		}
	}
}

// isBlockReached checks if the current block number is greater than or equal to the given one.
func (ex *Executor) isBlockReached(blockNumber uint64) (bool, error) {
	b, err := ex.client.BlockNumber(ex.ctx)
	if err != nil {
		return false, fmt.Errorf("error querying current block number: %w", err)
	}
	return b >= blockNumber, nil
}
