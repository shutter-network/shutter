package app

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

// Visit https://github.com/tendermint/spec/blob/master/spec/abci/abci.md for more information on
// the application interface we're implementing here.
// https://docs.tendermint.com/master/spec/abci/apps.html also provides some useful information

// CheckTx checks if a transaction is valid. If return Code != 0, it will be rejected from the
// mempool and hence not broadcasted to other peers and not included in a proposal block.
func (app *ShutterApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: 0, GasWanted: 1}
}

// NewShutterApp creates a new ShutterApp
func NewShutterApp() *ShutterApp {
	return &ShutterApp{
		Configs:     []*BatchConfig{{}},
		BatchStates: make(map[uint64]BatchState),
	}
}

// getConfig returns the BatchConfig for the given batchIndex
func (app *ShutterApp) getConfig(batchIndex uint64) *BatchConfig {
	for i := len(app.Configs) - 1; i >= 0; i-- {
		if app.Configs[i].StartBatchIndex <= batchIndex {
			return app.Configs[i]
		}
	}
	panic("guard element missing")
}

func (app *ShutterApp) addConfig(cfg BatchConfig) error {
	lastConfig := app.Configs[len(app.Configs)-1]
	if cfg.StartBatchIndex < lastConfig.StartBatchIndex {
		return fmt.Errorf(
			"start batch index of next config (%d) lower than current one (%d)",
			cfg.StartBatchIndex,
			lastConfig.StartBatchIndex,
		)
	}
	app.Configs = append(app.Configs, &cfg)
	return nil
}

// getBatchState returns the BatchState for the given batchIndex
func (app *ShutterApp) getBatchState(batchIndex uint64) BatchState {
	bs, ok := app.BatchStates[batchIndex]
	if !ok {
		bs.BatchIndex = batchIndex
		bs.Config = app.getConfig(batchIndex)
	}

	return bs
}

func (ShutterApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (ShutterApp) SetOption(req abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{}
}

/* BlockExecution

The first time a new blockchain is started, Tendermint calls InitChain. From then on, the following
sequence of methods is executed for each block:

BeginBlock, [DeliverTx], EndBlock, Commit

where one DeliverTx is called for each transaction in the block. The result is an updated
application state. Cryptographic commitments to the results of DeliverTx, EndBlock, and Commit are
included in the header of the next block.
*/
func (ShutterApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (ShutterApp) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

// decodeTx decodes the given transaction.  It's kind of strange that we have do URL decode the
// message outselves instead of tendermint doing it for us.
func (ShutterApp) decodeTx(req abcitypes.RequestDeliverTx) (signer common.Address, msg *shmsg.Message, err error) {
	var signedMsg []byte
	signedMsg, err = base64.RawURLEncoding.DecodeString(string(req.Tx))
	if err != nil {
		return
	}
	signer, err = shmsg.GetSigner(signedMsg)
	if err != nil {
		return
	}

	msg, err = shmsg.GetMessage(signedMsg)

	if err != nil {
		return
	}
	return
}

func (app *ShutterApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	signer, msg, err := app.decodeTx(req)
	if err != nil {
		fmt.Println("Error while decoding transaction:", err)
		return makeErrorResponse(fmt.Sprintf("Error while decoding transaction: %s", err))
	}
	return app.deliverMessage(msg, signer)
}

func makeErrorResponse(msg string) abcitypes.ResponseDeliverTx {
	return abcitypes.ResponseDeliverTx{
		Code:   1,
		Log:    msg,
		Events: []abcitypes.Event{},
	}
}

func (app *ShutterApp) deliverPublicKeyCommitment(pkc *shmsg.PublicKeyCommitment, sender common.Address) abcitypes.ResponseDeliverTx {
	bs := app.getBatchState(pkc.BatchIndex)
	publicKeyBefore := bs.PublicKey
	err := bs.AddPublicKeyCommitment(PublicKeyCommitment{Sender: sender, Pubkey: pkc.Commitment})
	if err != nil {
		fmt.Println("GOT ERROR", err)
		return makeErrorResponse(fmt.Sprintf("Error in AddPublicKeyCommitment: %s", err))
	}
	app.BatchStates[pkc.BatchIndex] = bs

	var events []abcitypes.Event
	if publicKeyBefore == nil && bs.PublicKey != nil {
		// we have generated a public key with this PublicKeyCommitment
		events = append(events, MakePubkeyGeneratedEvent(pkc.BatchIndex, bs.PublicKey))
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events,
	}
}

func (app *ShutterApp) deliverSecretShare(ss *shmsg.SecretShare, sender common.Address) abcitypes.ResponseDeliverTx {
	bs := app.getBatchState(ss.BatchIndex)
	privateKeyBefore := bs.PrivateKey
	err := bs.AddSecretShare(SecretShare{Sender: sender, Privkey: ss.Privkey})
	if err != nil {
		fmt.Println("GOT ERROR", err)
		return makeErrorResponse(fmt.Sprintf("Error in AddSecretShare: %s", err))
	}
	app.BatchStates[ss.BatchIndex] = bs
	var events []abcitypes.Event
	if privateKeyBefore == nil && bs.PrivateKey != nil {
		// we have generated a public key with this PublicKeyCommitment
		events = append(events, MakePrivkeyGeneratedEvent(ss.BatchIndex, bs.PrivateKey))
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events,
	}
}

func (app *ShutterApp) deliverBatchConfig(msg *shmsg.BatchConfig, sender common.Address) abcitypes.ResponseDeliverTx {
	// XXX everyone can call this at the moment
	var keypers []common.Address
	for _, k := range msg.Keypers {
		keypers = append(keypers, common.BytesToAddress(k))
	}

	bc := BatchConfig{
		StartBatchIndex: msg.StartBatchIndex,
		Keypers:         keypers,
		Threshold:       msg.Threshold,
	}
	err := app.addConfig(bc)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf("Error in addConfig: %s", err))
	}

	var events []abcitypes.Event
	events = append(events, MakeBatchConfigEvent(bc.StartBatchIndex, bc.Threshold, keypers))
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events,
	}
}

func (app *ShutterApp) deliverEncryptionKeyAttestation(msg *shmsg.EncryptionKeyAttestation, sender common.Address) abcitypes.ResponseDeliverTx {
	bs := app.getBatchState(msg.BatchIndex)
	att := EncryptionKeyAttestation{
		Sender:                sender,
		EncryptionKey:         msg.Key,
		BatchIndex:            msg.BatchIndex,
		ConfigContractAddress: common.BytesToAddress(msg.ConfigContractAddress),
		Signature:             msg.Signature,
	}
	err := bs.AddEncryptionKeyAttestation(att)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf("Error in AddEncryptionKeyAttestation: %s", err))
	}
	app.BatchStates[msg.BatchIndex] = bs

	keyperIndex, ok := bs.Config.KeyperIndex(sender)
	if !ok {
		// this is already checked in AddEncryptionKeyAttestation, but no harm in handling it twice
		return makeErrorResponse("not a keyper")
	}

	event := MakeEncryptionKeySignatureAddedEvent(keyperIndex, msg.BatchIndex, msg.Key, msg.Signature)
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{event},
	}
}

func (app *ShutterApp) deliverMessage(msg *shmsg.Message, sender common.Address) abcitypes.ResponseDeliverTx {
	fmt.Println("MSG:", msg)
	if msg.GetPublicKeyCommitment() != nil {
		return app.deliverPublicKeyCommitment(msg.GetPublicKeyCommitment(), sender)
	}
	if msg.GetSecretShare() != nil {
		return app.deliverSecretShare(msg.GetSecretShare(), sender)
	}
	if msg.GetBatchConfig() != nil {
		return app.deliverBatchConfig(msg.GetBatchConfig(), sender)
	}
	if msg.GetEncryptionKeyAttestation() != nil {
		return app.deliverEncryptionKeyAttestation(msg.GetEncryptionKeyAttestation(), sender)
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{},
	}
}

func (ShutterApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *ShutterApp) Commit() abcitypes.ResponseCommit {
	return abcitypes.ResponseCommit{}
}
