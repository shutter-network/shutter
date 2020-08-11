package app

import (
	"github.com/tendermint/tendermint/abci/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

// Visit https://github.com/tendermint/spec/blob/master/spec/abci/abci.md for more information on
// the application interface we're implementing here.
// https://docs.tendermint.com/master/spec/abci/apps.html also provides some useful information

type KeyGenerationRound struct {
	votes []string
}

type ShutterApp struct {
}

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

func (app *ShutterApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
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
