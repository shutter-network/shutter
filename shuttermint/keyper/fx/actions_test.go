package fx

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"

	"github.com/shutter-network/shutter/shlib/shcrypto"
	"github.com/shutter-network/shutter/shlib/shtest"
	"github.com/shutter-network/shutter/shuttermint/medley"
	"github.com/shutter-network/shutter/shuttermint/shmsg"
)

// exampleStruct is a struct which contains an IAction valued field, in order to test serializing
// IActions with gob.
type exampleStruct struct {
	Action IAction
}

func sendShuttermintMessage() *SendShuttermintMessage {
	keypers := []common.Address{common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2))}
	msg := shmsg.NewBatchConfig(
		200,
		keypers,
		2,
		common.BigToAddress(big.NewInt(1024)),
		55,
		false,
		false,
	)
	return &SendShuttermintMessage{
		Description: "foo bar baz",
		Msg:         msg,
	}
}

// TestSendShuttermintMessageGobable tests that a SendShuttermintMessage is gobable. We need a
// dedicated test here, because SendShuttermintMessage contains a protocol buffer Msg field.
func TestSendShuttermintMessageGobable(t *testing.T) {
	a1 := exampleStruct{Action: sendShuttermintMessage()}
	a2 := exampleStruct{}

	medley.CloneWithGob(a1, &a2)
	s1 := a1.Action.(*SendShuttermintMessage)
	s2 := a2.Action.(*SendShuttermintMessage)
	assert.Equal(t, s1.Description, s2.Description)
	assert.Assert(t, proto.Equal(s1.Msg, s2.Msg))
}

var actions []IAction = []IAction{
	&ExecuteCipherBatch{BatchIndex: 55, KeyperIndex: 11},
	&ExecutePlainBatch{BatchIndex: 56},
	&SkipCipherBatch{BatchIndex: 57},
	&Accuse{HalfStep: 55},
	&Appeal{},
	&EonKeyBroadcast{
		KeyperIndex:     11,
		StartBatchIndex: 12,
		EonPublicKey:    (*shcrypto.EonPublicKey)(new(bn256.G2).ScalarBaseMult(big.NewInt(5))),
	},
}

func TestActionsGobable(t *testing.T) {
	for i, a := range actions {
		a1 := exampleStruct{Action: a}
		a2 := exampleStruct{}
		t.Run(fmt.Sprintf("%d: %s", i, a), func(t *testing.T) {
			shtest.EnsureGobable(t, &a1, &a2)
		})
	}
}
