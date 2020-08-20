package app

import (
	"encoding/base64"
	"fmt"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/abci/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// Visit https://github.com/tendermint/spec/blob/master/spec/abci/abci.md for more information on
// the application interface we're implementing here.
// https://docs.tendermint.com/master/spec/abci/apps.html also provides some useful information

var _ abcitypes.Application = (*ShutterApp)(nil)

// CheckTx checks if a transaction is valid. If return Code != 0, it will be rejected from the
// mempool and hence not broadcasted to other peers and not included in a proposal block.

func (app *ShutterApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {

	return abcitypes.ResponseCheckTx{Code: 0, GasWanted: 1}
}

func NewShutterApp() *ShutterApp {
	app := ShutterApp{}
	return &app
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
	// fmt.Println("BEGIN", req.GetHash(), req.Header.GetChainID(), req.Header.GetTime())
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

	// XXX we need to check that the signer is allowed
	fmt.Println("Signer:", signer.Hex())

	msg, err = shmsg.GetMessage(signedMsg)

	if err != nil {
		return
	}
	return
}

func (app *ShutterApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	// signedTransaction := make([]byte,  base64.RawURLEncoding.DecodedLen(len(req.Tx)))

	signer, msg, err := app.decodeTx(req)

	if err != nil {
		fmt.Println("Error while decoding transaction:", err)
		return abcitypes.ResponseDeliverTx{
			Code:   1,
			Events: []types.Event{}}
	}
	return app.deliverMessage(msg, signer)
}

func (app *ShutterApp) deliverPublicKeyCommitment(pkc *shmsg.PublicKeyCommitment, sender common.Address) abcitypes.ResponseDeliverTx {
	bk, _ := app.Batches[pkc.BatchIndex]
	err := bk.AddPublicKeyCommitment(PublicKeyCommitment{Sender: sender, Pubkey: pkc.Commitment})
	if err != nil {
		return abcitypes.ResponseDeliverTx{
			Code:   1,
			Events: []types.Event{}}

	}
	app.Batches[pkc.BatchIndex] = bk

	var events []types.Event
	if len(bk.Commitments) == int(bk.Config.Threshhold) {
		// we have reached the threshold
		// for the fake key generation let's take the last key as public key for this batch
		events = append(events, types.Event{
			Type: "shutter.pubkey-generated",
			Attributes: []kv.Pair{
				{Key: []byte("BatchIndex"), Value: []byte("XXX")},
				{Key: []byte("Pubkey"), Value: []byte("XXX")}},
		})
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events}
}

func (app *ShutterApp) deliverMessage(msg *shmsg.Message, sender common.Address) abcitypes.ResponseDeliverTx {
	fmt.Println("MSG:", msg)
	if msg.GetPublicKeyCommitment() != nil {
		return app.deliverPublicKeyCommitment(msg.GetPublicKeyCommitment(), sender)
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []types.Event{}}
}

func (ShutterApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	// fmt.Println("END BLOCK", req.Height)
	return abcitypes.ResponseEndBlock{}
}

func (app *ShutterApp) Commit() abcitypes.ResponseCommit {
	// fmt.Printf("COMMIT %#v", app.games)
	return abcitypes.ResponseCommit{}
}

func (ShutterApp) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{Code: 0}
}
