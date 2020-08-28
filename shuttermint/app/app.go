package app

import (
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/abci/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/kv"
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
		Configs: []*BatchConfig{{}},
		Batches: make(map[uint64]BatchKeys),
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
	if lastConfig.StartBatchIndex >= cfg.StartBatchIndex {
		return errors.New("StartBatchIndex must be greater than previous StartBatchIndex")
	}
	app.Configs = append(app.Configs, &cfg)
	return nil
}

// getBatch returns the BatchKeys for the given batchIndex
func (app *ShutterApp) getBatch(batchIndex uint64) BatchKeys {
	bk, ok := app.Batches[batchIndex]
	if !ok {
		bk.Config = app.getConfig(batchIndex)
	}

	return bk
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

// encodePubkeyForEvent encodes the PublicKey as a string suitable for putting it into a tendermint
// event, i.e. an utf-8 compatible string
func encodePubkeyForEvent(pubkey *ecdsa.PublicKey) string {
	return base64.RawURLEncoding.EncodeToString(crypto.FromECDSAPub(pubkey))
}

func decodePubkeyFromEvent(s string) (*ecdsa.PublicKey, error) {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalPubkey(data)
}

func encodePrivkeyForEvent(privkey *ecdsa.PrivateKey) string {
	return base64.RawURLEncoding.EncodeToString(crypto.FromECDSA(privkey))
}

func (app *ShutterApp) deliverPublicKeyCommitment(pkc *shmsg.PublicKeyCommitment, sender common.Address) abcitypes.ResponseDeliverTx {
	bk := app.getBatch(pkc.BatchIndex)
	publicKeyBefore := bk.PublicKey
	err := bk.AddPublicKeyCommitment(PublicKeyCommitment{Sender: sender, Pubkey: pkc.Commitment})
	if err != nil {
		fmt.Println("GOT ERROR", err)
		return abcitypes.ResponseDeliverTx{
			Code:   1,
			Events: []types.Event{}}

	}
	app.Batches[pkc.BatchIndex] = bk

	var events []types.Event
	if publicKeyBefore == nil && bk.PublicKey != nil {
		// we have generated a public key with this PublicKeyCommitment
		events = append(events, types.Event{
			Type: "shutter.pubkey-generated",
			Attributes: []kv.Pair{
				{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", pkc.BatchIndex))},
				{Key: []byte("Pubkey"), Value: []byte(encodePubkeyForEvent(bk.PublicKey))}},
		})
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events}
}

func (app *ShutterApp) deliverSecretShare(ss *shmsg.SecretShare, sender common.Address) abcitypes.ResponseDeliverTx {
	bk := app.getBatch(ss.BatchIndex)
	privateKeyBefore := bk.PrivateKey
	err := bk.AddSecretShare(SecretShare{Sender: sender, Privkey: ss.Privkey})
	if err != nil {
		fmt.Println("GOT ERROR", err)
		return abcitypes.ResponseDeliverTx{
			Code:   1,
			Events: []types.Event{}}
	}
	app.Batches[ss.BatchIndex] = bk
	var events []types.Event
	if privateKeyBefore == nil && bk.PrivateKey != nil {
		// we have generated a public key with this PublicKeyCommitment
		events = append(events, types.Event{
			Type: "shutter.privkey-generated",
			Attributes: []kv.Pair{
				{Key: []byte("BatchIndex"), Value: []byte(fmt.Sprintf("%d", ss.BatchIndex))},
				{Key: []byte("Privkey"), Value: []byte(encodePrivkeyForEvent(bk.PrivateKey))}},
		})
	}
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events}

}

func encodeAddressesForEvent(addr []common.Address) string {
	var hex []string
	for _, a := range addr {
		hex = append(hex, a.Hex())
	}
	return strings.Join(hex, ",")
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
		Threshhold:      msg.Threshold,
	}
	err := app.addConfig(bc)
	if err != nil {
		return abcitypes.ResponseDeliverTx{
			Code:   1,
			Events: []types.Event{}}
	}

	var events []types.Event
	events = append(events, types.Event{
		Type: "shutter.batch-config",
		Attributes: []kv.Pair{
			{Key: []byte("StartBatchIndex"), Value: []byte(fmt.Sprintf("%d", bc.StartBatchIndex))},
			{Key: []byte("Threshhold"), Value: []byte(fmt.Sprintf("%d", bc.Threshhold))},
			{Key: []byte("Keypers"), Value: []byte(encodeAddressesForEvent(keypers))},
		},
	})
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events}
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
