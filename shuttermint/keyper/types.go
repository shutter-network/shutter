package keyper

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"
)

// RoundInterval is the duration between the start of two consecutive rounds
var RoundInterval time.Duration = time.Duration(5 * time.Second)

// PrivateKeyDeley is the duration between the start of the public key generation and the the start
// of the private key generation for a single round
var PrivateKeyDelay time.Duration = time.Duration(45 * time.Second)

type BatchParams struct {
	BatchIndex                   uint64
	PublicKeyGenerationStartTime time.Time
	// PublicKeyGenerationDuration  time.Duration
	PrivateKeyGenerationStartTime time.Time
}

func NewBatchParams(BatchIndex uint64) BatchParams {
	ts := int64(BatchIndex) * int64(RoundInterval)

	pubstart := time.Unix(ts/int64(time.Second), ts%int64(time.Second))
	privstart := pubstart.Add(PrivateKeyDelay)
	return BatchParams{
		BatchIndex:                    BatchIndex,
		PublicKeyGenerationStartTime:  pubstart,
		PrivateKeyGenerationStartTime: privstart,
	}
}

func NextBatchIndex(t time.Time) uint64 {
	return uint64((t.UnixNano() + int64(RoundInterval) - 1) / int64(RoundInterval))
}

type BatchRunner struct {
	rpcclient client.Client
	params    BatchParams
}

type MessageSender struct {
	rpcclient  client.Client
	signingKey *ecdsa.PrivateKey
}

func NewMessageSender(client client.Client, signingKey *ecdsa.PrivateKey) MessageSender {
	return MessageSender{client, signingKey}
}

func (ms MessageSender) SendMessage(msg *shmsg.Message) error {
	signedMessage, err := shmsg.SignMessage(msg, ms.signingKey)
	if err != nil {
		return err
	}
	var tx types.Tx = types.Tx(base64.RawURLEncoding.EncodeToString(signedMessage))
	res, err := ms.rpcclient.BroadcastTxCommit(tx)
	if err != nil {
		return err
	}
	// fmt.Println("broadcast tx", res)
	if res.DeliverTx.Code != 0 {
		return fmt.Errorf("Error in SendMessage: %s", res.DeliverTx.Log)
	}
	return nil
}
