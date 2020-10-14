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
			log.Printf("Fast-forwarded to batch #%d", p.BatchIndex)

			// only execute cipher batch here, plain batch will be executed during next fast
			// forward
			err = ex.executeCipherBatch(p)
			if err != nil {
				log.Printf("Error executing batch #%d: %s", p.BatchIndex, err)
			}
			log.Printf("Executed cipher batch #%d", p.BatchIndex)

		case <-time.After(skipCheckInterval):
			err := ex.fastForward(math.MaxUint64, false)
			if err != nil {
				log.Printf("Error skipping batches: %s", err)
			}

		case <-ex.ctx.Done():
			return nil
		}
	}
}

// executeBatch executes the current cipher batch using the given parameters.
func (ex *Executor) executeCipherBatch(p CipherExecutionParams) error {
	kickOffBlock, err := ex.cipherExecutionKickOffBlock(p.BatchIndex)
	if err != nil {
		return err
	}
	blockReached, err := ex.waitForBlockOrHalfStep(kickOffBlock, p.CipherHalfStep()+1)
	if err != nil {
		return err
	}
	if !blockReached {
		return nil
	}

	// check if we're (still) at the right half step. If not, someone else has probably executed
	// it while we were waiting for the kick off block.
	numHalfSteps, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(ex.callOpts(nil))
	if err != nil {
		return err
	}
	isCipher := numHalfSteps%2 == 0
	batchIndex := numHalfSteps / 2
	if !isCipher || batchIndex != p.BatchIndex {
		return nil
	}

	return ex.executeCipherHalfStep(p)
}

// fastForward skips all cipher batches that are skippable right now and executes all plain
// batches. It will stop when lastBatchIndex has been handled. If wait is false, it will return as
// soon as immediate progress cannot be made. Otherwise, it will wait.
func (ex *Executor) fastForward(lastBatchIndex uint64, wait bool) error {
	for {
		numHalfSteps, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(ex.callOpts(nil))
		if err != nil {
			return err
		}
		isCipher := numHalfSteps%2 == 0
		batchIndex := numHalfSteps / 2

		if batchIndex > lastBatchIndex {
			return nil
		}

		var kickOffBlock uint64
		if isCipher {
			kickOffBlock, err = ex.cipherSkipKickOffBlock(batchIndex)
		} else {
			kickOffBlock, err = ex.plainExecutionKickOffBlock(batchIndex)
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
			err = ex.skipCipherHalfStep(batchIndex)
		} else {
			err = ex.executePlainHalfStep(batchIndex)
		}
		if err != nil {
			return err
		}
	}
}

// executePlainHalfStep executes the plain portion of the current batch and returns when it's
// done. It does not perform any prior checks.
func (ex *Executor) executePlainHalfStep(batchIndex uint64) error {
	tx, err := ex.cc.ExecutorContract.ExecutePlainBatch(ex.transactOpts(), [][]byte{})
	if err != nil {
		return fmt.Errorf("error sending plain batch execution tx: %w", err)
	}
	log.Printf("Plain batch execution tx sent for batch #%d: %s", batchIndex, tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, (batchIndex+1)*2)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %w", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to execute plain half step", tx.Hash().Hex())
		}
		log.Printf("Plain batch execution tx for batch #%d successful", batchIndex)
	} else {
		log.Printf("Plain batch execution tx for batch #%d potentially redundant", batchIndex)
	}

	return nil
}

// skipCipherHalfStep skips the current cipher half step and returns when it's done. It does not
// perform any prior checks.
func (ex *Executor) skipCipherHalfStep(batchIndex uint64) error {
	tx, err := ex.cc.ExecutorContract.SkipCipherExecution(ex.transactOpts())
	if err != nil {
		return fmt.Errorf("error sending skip cipher execution tx: %w", err)
	}
	log.Printf("Cipher execution skip tx sent for batch #%d: %s", batchIndex, tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, batchIndex*2+1)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %w", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to skip cipher half step", tx.Hash().Hex())
		}
		log.Printf("cipher half step successfully skipped in tx %s", tx.Hash().Hex())
	} else {
		log.Printf("cipher half step skipped by another keyper")
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
	log.Printf("Cipher batch execution tx sent for batch #%d: %s", p.BatchIndex, tx.Hash().Hex())

	receipt, err := ex.waitForReceiptOrHalfStep(tx, p.CipherHalfStep()+1)
	if err != nil {
		return fmt.Errorf("error waiting for receipt for tx %s: %s", tx.Hash().Hex(), err)
	}
	if receipt != nil {
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("tx %s has failed to execute cipher half step", tx.Hash().Hex())
		}
		log.Printf("cipher half step successfully executed in tx %s", tx.Hash().Hex())
	} else {
		log.Printf("cipher half step executed by another keyper")
	}

	return nil
}

// cipherSkipKickOffBlock returns the block number at which the cipher half step should be
// skipped.
func (ex *Executor) cipherSkipKickOffBlock(batchIndex uint64) (uint64, error) {
	batchParams, err := ex.cc.ConfigContract.QueryBatchParams(ex.callOpts(nil), batchIndex)
	if err != nil {
		return 0, err
	}

	return batchParams.ExecutionTimeoutBlock, nil
}

// plainExecutionKickOffBlock returns the block number at which the plain half step should be
// executed.
func (ex *Executor) plainExecutionKickOffBlock(batchIndex uint64) (uint64, error) {
	return 0, nil
}

// cipherExecutionKickOffBlock returns the block number at which the cipher half step should be
// executed.
func (ex *Executor) cipherExecutionKickOffBlock(batchIndex uint64) (uint64, error) {
	return 0, nil
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

// headerChannel creates a channel to which the current and all future block headers will be sent.
func (ex *Executor) headerChannel(ctx context.Context) (<-chan *types.Header, <-chan error) {
	headers := make(chan *types.Header)
	errors := make(chan error)

	go func() {
		defer close(headers)
		defer close(errors)

		sub, err := ex.client.SubscribeNewHead(ctx, headers)
		if err != nil {
			errors <- err
			return
		}
		defer sub.Unsubscribe()

		firstHeader, err := ex.client.HeaderByNumber(ctx, nil)
		if err != nil {
			errors <- err
			return
		}
		headers <- firstHeader

		for {
			select {
			case <-ctx.Done():
				errors <- ctx.Err()
				return
			case err := <-sub.Err():
				errors <- err
				return
			case header := <-headers:
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
			errors <- err
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
			errors <- err
			return
		}
		callOpts := &bind.CallOpts{
			Context:     ctx,
			BlockNumber: header.Number,
		}
		startHalfStep, err := ex.cc.ExecutorContract.NumExecutionHalfSteps(callOpts)
		halfSteps <- startHalfStep

		// subscribe to all events that change half step
		startBlock := header.Number.Uint64()
		watchOpts := &bind.WatchOpts{
			Context: ctx,
			Start:   &startBlock,
		}

		batchExecutedEvents := make(chan *contract.ExecutorContractBatchExecuted)
		batchExecutedSub, err := ex.cc.ExecutorContract.WatchBatchExecuted(watchOpts, batchExecutedEvents)
		if err != nil {
			errors <- err
			return
		}
		defer batchExecutedSub.Unsubscribe()

		cipherSkippedEvents := make(chan *contract.ExecutorContractCipherExecutionSkipped)
		cipherSkippedSub, err := ex.cc.ExecutorContract.WatchCipherExecutionSkipped(watchOpts, cipherSkippedEvents)
		if err != nil {
			errors <- err
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
				errors <- err
				return
			case err := <-cipherSkippedSub.Err():
				errors <- err
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
			return nil, fmt.Errorf("error waiting for receipt: %w", err)
		case err := <-halfStepErrors:
			return nil, fmt.Errorf("error waiting for half step: %w", err)
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
		return false, err
	}
	return b >= blockNumber, nil
}
