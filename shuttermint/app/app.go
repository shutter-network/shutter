package app

import (
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/go-amino"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

var (
	// PersistMinDuration is the minimum duration between two calls to persistToDisk
	// TODO we should probably increase the default here and we should have a way to collect
	// garbage to keep the persisted state small enough.
	// The variable is declared here, because we do not want to persist it as part of the
	// application. The same could be said about the Gobpath field though, which we persist as
	// part of the application.
	// If we set this to zero, the state will get saved on every call to Commit
	PersistMinDuration time.Duration = 30 * time.Second

	// NonExistentValidator is an artificial key used to replace the voting power of validators
	// that haven't sent their CheckIn message yet
	NonExistentValidator ValidatorPubkey
)

func init() {
	gob.Register(crypto.S256()) // Allow gob to serialize ecsda.PrivateKey

	var err error
	k := [32]byte{'n', 'o', 'v', 'a', 'l', 'i', 'd', 'a', 't', 'o', 'r'}
	NonExistentValidator, err = NewValidatorPubkey(k[:])
	if err != nil {
		panic(err)
	}
}

// ABCIPubkey returns an abcitypes.PubKey struct suitable for updating the shuttermint validators
func (vp ValidatorPubkey) ABCIPubkey() abcitypes.PubKey {
	return abcitypes.PubKey{Type: "ed25519", Data: []byte(vp.Ed25519pubkey)}
}

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
		Configs:      []*BatchConfig{{}},
		BatchStates:  make(map[uint64]BatchState),
		Voting:       NewConfigVoting(),
		Identities:   make(map[common.Address]ValidatorPubkey),
		StartedVotes: make(map[common.Address]bool),
	}
}

// LoadShutterAppFromFile loads a shutter app from a file
func LoadShutterAppFromFile(gobpath string) (ShutterApp, error) {
	var shapp ShutterApp
	gobfile, err := os.Open(gobpath)
	if os.IsNotExist(err) {
		shapp = *NewShutterApp()
	} else if err != nil {
		return shapp, err
	} else {
		defer gobfile.Close()
		dec := gob.NewDecoder(gobfile)
		err = dec.Decode(&shapp)
		if err != nil {
			return shapp, err
		}
		log.Printf("Loaded shutter app from file %s, last saved %s, block height %d",
			gobpath, shapp.LastSaved, shapp.LastBlockHeight)
	}

	shapp.Gobpath = gobpath
	shapp.LastSaved = time.Now() // Do not persist immediately after starting
	return shapp, nil
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

// checkConfig checks if the given BatchConfig could be added.
func (app *ShutterApp) checkConfig(cfg BatchConfig) error {
	err := cfg.EnsureValid()
	if err != nil {
		return err
	}
	lastConfig := app.LastConfig()
	if cfg.StartBatchIndex < lastConfig.StartBatchIndex {
		return fmt.Errorf(
			"start batch index of next config (%d) lower than current one (%d)",
			cfg.StartBatchIndex,
			lastConfig.StartBatchIndex,
		)
	}
	return nil
}

func (app *ShutterApp) addConfig(cfg BatchConfig) error {
	err := app.checkConfig(cfg)
	if err != nil {
		return err
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

// Info should return the latest committed state of the app. On startup, tendermint calls the Info
// method and will replay blocks that are not yet committed.
// See https://github.com/tendermint/spec/blob/master/spec/abci/apps.md#crash-recovery
func (app *ShutterApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{
		LastBlockHeight:  app.LastBlockHeight,
		LastBlockAppHash: []byte(""),
	}
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
func (app *ShutterApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	genesisState := GenesisAppState{}
	err := amino.NewCodec().UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		log.Fatalf("Cannot unmarshal genesis app state: %s", err)
	}

	bc := BatchConfig{
		StartBatchIndex: 0,
		Keypers:         genesisState.GetKeypers(),
		Threshold:       genesisState.Threshold,
	}
	err = bc.EnsureValid()
	if err != nil {
		log.Fatalf("Invalid genesis app state: %s", err)
	}

	if len(app.Configs) == 1 && len(app.Configs[0].Keypers) == 0 {
		log.Print("Initializing new chain")
		for i, k := range genesisState.Keypers {
			log.Printf("Initial keyper #%d: %s", i, k.String())
		}
		validators, err := MakePowermap(req.Validators)
		if err != nil {
			log.Fatalf("InitChain: cannot handle validator keys: %s", err)
		}
		app.Validators = validators
		app.Configs = []*BatchConfig{&bc}
	} else if !reflect.DeepEqual(bc, *app.Configs[0]) {
		log.Fatalf("Mismatch between stored app state and initial app state, stored=%+v initial=%+v", app.Configs[0], bc)
	}
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
		msg := fmt.Sprintf("Error while decoding transaction: %s", err)
		log.Print(msg)
		return makeErrorResponse(msg)
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

func notAKeyper(sender common.Address) abcitypes.ResponseDeliverTx {
	return makeErrorResponse(fmt.Sprintf(
		"sender %s is not a keyper", sender.Hex()))
}

func (app *ShutterApp) deliverPublicKeyCommitment(pkc *shmsg.PublicKeyCommitment, sender common.Address) abcitypes.ResponseDeliverTx {
	bs := app.getBatchState(pkc.BatchIndex)
	publicKeyBefore := bs.PublicKey
	err := bs.AddPublicKeyCommitment(PublicKeyCommitment{Sender: sender, Pubkey: pkc.Commitment})
	if err != nil {
		msg := fmt.Sprintf("Error: cannot deliver public key commitment: %s", err)
		log.Print(msg)
		return makeErrorResponse(msg)
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
		msg := fmt.Sprintf("Error: cannot add secret share: %s", err)
		log.Print(msg)
		return makeErrorResponse(msg)
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

func (app *ShutterApp) allowedToVoteOnConfigChanges(sender common.Address) bool {
	lastConfig := app.LastConfig()
	_, ok := lastConfig.KeyperIndex(sender)
	return ok
}

func (app *ShutterApp) deliverBatchConfig(msg *shmsg.BatchConfig, sender common.Address) abcitypes.ResponseDeliverTx {
	bc, err := BatchConfigFromMessage(msg)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf("Malformed BatchConfig message: %s", err))
	}

	err = app.checkConfig(bc)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf("checkConfig: %s", err))
	}

	if !app.allowedToVoteOnConfigChanges(sender) {
		return makeErrorResponse("not allowed to vote on config changes")
	}

	var events []abcitypes.Event

	err = app.Voting.AddVote(sender, bc)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf("Error in addConfig: %s", err))
	}

	_, ok := app.Voting.Outcome(int(app.LastConfig().Threshold))
	if ok {
		app.Voting = NewConfigVoting()
		err = app.addConfig(bc)
		if err != nil {
			return makeErrorResponse(fmt.Sprintf("Error in addConfig: %s", err))
		}

		events = append(events, MakeBatchConfigEvent(bc.StartBatchIndex, bc.Threshold, bc.Keypers))
	}

	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: events,
	}
}

func (app *ShutterApp) deliverEncryptionKeyAttestation(
	msg *shmsg.EncryptionKeyAttestation,
	sender common.Address,
) abcitypes.ResponseDeliverTx {
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
		return notAKeyper(sender)
	}

	event := MakeEncryptionKeySignatureAddedEvent(keyperIndex, msg.BatchIndex, msg.Key, msg.Signature)
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{event},
	}
}

// isKeyper checks if the given address is a keyper in any config (current and previous ones)
func (app *ShutterApp) isKeyper(a common.Address) bool {
	for _, cfg := range app.Configs {
		_, ok := cfg.KeyperIndex(a)
		if ok {
			return true
		}
	}
	return false
}

func (app *ShutterApp) deliverCheckIn(msg *shmsg.CheckIn, sender common.Address) abcitypes.ResponseDeliverTx {
	_, ok := app.Identities[sender]
	if ok {
		return makeErrorResponse(fmt.Sprintf(
			"sender %s already checked in", sender.Hex()))
	}
	if !app.isKeyper(sender) {
		return notAKeyper(sender)
	}

	pk, err := NewValidatorPubkey(msg.Pubkey)
	if err != nil {
		return makeErrorResponse(fmt.Sprintf(
			"malformed pubkey: %s", err))
	}
	app.Identities[sender] = pk
	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{},
	}
}

func (app *ShutterApp) deliverBatchConfigStarted(msg *shmsg.BatchConfigStarted, sender common.Address) abcitypes.ResponseDeliverTx {
	configIndex := msg.GetBatchConfigIndex()
	lastBatchConfig := app.LastConfig()
	if configIndex != lastBatchConfig.ConfigIndex {
		return makeErrorResponse(fmt.Sprintf(
			"can only start last config with index %d, got index %d",
			lastBatchConfig.ConfigIndex, configIndex))
	}
	if len(app.Configs) <= 1 {
		return makeErrorResponse("no config to vote on")
	}
	// We have to look into the config before the last config to see if we're allowed to vote
	if !app.Configs[len(app.Configs)-2].IsKeyper(sender) {
		return notAKeyper(sender)
	}

	if !lastBatchConfig.Started {
		app.StartedVotes[sender] = true
	}

	return abcitypes.ResponseDeliverTx{
		Code:   0,
		Events: []abcitypes.Event{},
	}
}

func (app *ShutterApp) deliverMessage(msg *shmsg.Message, sender common.Address) abcitypes.ResponseDeliverTx {
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
	if msg.GetBatchConfigStarted() != nil {
		return app.deliverBatchConfigStarted(msg.GetBatchConfigStarted(), sender)
	}
	if msg.GetCheckIn() != nil {
		return app.deliverCheckIn(msg.GetCheckIn(), sender)
	}
	log.Print("Error: cannot deliver messsage", msg)
	return makeErrorResponse("cannot deliver message")
}

func (app *ShutterApp) LastConfig() *BatchConfig {
	if len(app.Configs) == 0 {
		panic("internal error: app.Configs is empty")
	}
	return app.Configs[len(app.Configs)-1]
}

// makePowermap creates a power map for the given slice of keypers. The voting power of each keyper
// that hasn't registered yet, is given to the NonExistentValidator key
func (app *ShutterApp) makePowermap(keypers []common.Address) Powermap {
	pm := make(Powermap)
	for _, k := range keypers {
		pk, ok := app.Identities[k]
		if ok {
			pm[pk] += 10
		} else {
			pm[NonExistentValidator] += 10
		}
	}
	return pm
}

// CurrentValidators returns a powermap of current validators.
func (app *ShutterApp) CurrentValidators() Powermap {
	for i := len(app.Configs) - 1; i >= 0; i-- {
		if app.Configs[i].Started && app.Configs[i].ValidatorsUpdated {
			return app.makePowermap(app.Configs[i].Keypers)
		}
	}
	return app.Validators
}

// countCheckedInKeypers counts the number of keypers that have already checked in in the given slice
func (app *ShutterApp) countCheckedInKeypers(keypers []common.Address) uint64 {
	var numCheckedIn uint64
	for _, k := range keypers {
		_, ok := app.Identities[k]
		if ok {
			numCheckedIn++
		}
	}
	return numCheckedIn
}

func (app *ShutterApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	lastConfig := app.LastConfig()

	// start last config if there are enough votes
	if len(app.Configs) >= 2 {
		currentConfig := app.Configs[len(app.Configs)-2]
		if uint64(len(app.StartedVotes)) >= currentConfig.Threshold {
			log.Printf("starting config %d", lastConfig.ConfigIndex)
			lastConfig.Started = true
			app.StartedVotes = make(map[common.Address]bool)
		}
	}

	if lastConfig.Started && !lastConfig.ValidatorsUpdated && app.countCheckedInKeypers(lastConfig.Keypers) >= lastConfig.Threshold {
		lastConfig.ValidatorsUpdated = true
	}

	newValidators := app.CurrentValidators()
	validatorUpdates := DiffPowermaps(app.Validators, newValidators).ValidatorUpdates()
	app.Validators = newValidators
	app.LastBlockHeight = req.Height
	return abcitypes.ResponseEndBlock{ValidatorUpdates: validatorUpdates}
}

// persistToDisk stores the ShutterApp on disk. This method first writes to a temporary file and
// renames the file later. Most probably this will not work on windows!
func (app *ShutterApp) persistToDisk() error {
	log.Printf("Persisting state to disk, height=%d", app.LastBlockHeight)
	tmppath := app.Gobpath + ".tmp"
	file, err := os.Create(tmppath)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("Error: close file: %s", err)
			return
		}
	}()

	app.LastSaved = time.Now()
	enc := gob.NewEncoder(file)
	err = enc.Encode(app)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	err = os.Rename(tmppath, app.Gobpath)
	return err
}

func (app *ShutterApp) maybePersistToDisk() error {
	if app.Gobpath == "" {
		return nil
	}
	if time.Since(app.LastSaved) <= PersistMinDuration {
		return nil
	}
	return app.persistToDisk()
}

func (app *ShutterApp) Commit() abcitypes.ResponseCommit {
	err := app.maybePersistToDisk()
	if err != nil {
		log.Printf("Error: cannot persist state to disk: %s", err)
	}

	return abcitypes.ResponseCommit{}
}
