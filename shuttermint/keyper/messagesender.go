package keyper

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/tendermint/tendermint/rpc/client"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

var mockMessageSenderBufferSize = 0x10000

func init() {
	rand.Seed(int64(time.Now().UnixNano())) // Seed the PRNG we use for random nonces
}

// NewRPCMessageSender creates a new RPCMessageSender
func NewRPCMessageSender(cl client.Client, signingKey *ecdsa.PrivateKey) RPCMessageSender {
	return RPCMessageSender{cl, signingKey}
}

// SendMessage signs the given shmsg.Message and sends the message to shuttermint
func (ms RPCMessageSender) SendMessage(ctx context.Context, msg *shmsg.Message) error {
	msgWithNonce := ms.addNonce(msg)
	signedMessage, err := shmsg.SignMessage(msgWithNonce, ms.signingKey)
	if err != nil {
		return err
	}
	var tx tmtypes.Tx = tmtypes.Tx(base64.RawURLEncoding.EncodeToString(signedMessage))
	res, err := ms.rpcclient.BroadcastTxCommit(ctx, tx)
	if err != nil {
		return err
	}
	if res.DeliverTx.Code != 0 {
		return fmt.Errorf("remote error: %s", res.DeliverTx.Log)
	}
	return nil
}

func (ms RPCMessageSender) addNonce(msg *shmsg.Message) *shmsg.MessageWithNonce {
	return &shmsg.MessageWithNonce{
		RandomNonce: randomNonce(),
		Msg:         msg,
	}
}

func randomNonce() uint64 {
	return rand.Uint64()
}

// NewMockMessageSender creates a new MockMessageSender. We use a buffered channel with a rather
// large size in order to simplify writing our tests.
func NewMockMessageSender() MockMessageSender {
	return MockMessageSender{
		Msgs: make(chan *shmsg.Message, mockMessageSenderBufferSize),
	}
}

func (ms MockMessageSender) SendMessage(_ context.Context, msg *shmsg.Message) error {
	ms.Msgs <- msg
	return nil
}
