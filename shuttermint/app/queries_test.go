package app

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"

	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func TestQueryInvalidURL(t *testing.T) {
	invalidQueries := []string{
		"http://configs?batchIndex=0",
		"host/configs?batchIndex=0#fragment",
		"user@host/configs?batchIndex=0",
		"/configs?batchIndex=0#fragment",
		"/unknownmethod/",
	}

	app := NewShutterApp()
	for _, path := range invalidQueries {
		req := abcitypes.RequestQuery{
			Path: path,
		}
		res := app.Query(req)
		require.Equal(t, uint32(1), res.Code)
		t.Log(res.Log)
	}
}

func TestQueryInvalidConfig(t *testing.T) {
	invalidQueries := []string{
		"/configs",
		"/configs?batchIndex=-1",
		"/configs?batchIndex=one",
		"/configs?last=false",
		"/configs?batchIndex=1?last=true",
	}

	app := NewShutterApp()
	for _, path := range invalidQueries {
		req := abcitypes.RequestQuery{
			Path: path,
		}
		res := app.Query(req)
		require.Equal(t, uint32(1), res.Code)
		t.Log(res.Log)
	}
}

func TestQueryConfig(t *testing.T) {
	app := NewShutterApp()
	c1 := BatchConfig{
		ConfigIndex:     1,
		StartBatchIndex: 100,
		Threshold:       1,
		Keypers: []common.Address{
			common.BigToAddress(big.NewInt(0)),
			common.BigToAddress(big.NewInt(1)),
		},
	}
	c2 := BatchConfig{
		ConfigIndex:     2,
		StartBatchIndex: 200,
		Threshold:       1,
		Keypers: []common.Address{
			common.BigToAddress(big.NewInt(2)),
		},
	}
	for _, c := range []BatchConfig{c1, c2} {
		err := app.addConfig(c)
		require.Nil(t, err)
	}

	testCases := []struct {
		path   string
		config BatchConfig
	}{
		{
			path:   "/configs?batchIndex=100",
			config: c1,
		},
		{
			path:   "/configs?batchIndex=199",
			config: c1,
		},
		{
			path:   "/configs?batchIndex=200",
			config: c2,
		},
		{
			path:   "/configs?last=true",
			config: c2,
		},
	}

	for _, testCase := range testCases {
		req := abcitypes.RequestQuery{
			Path: testCase.path,
		}
		res := app.Query(req)
		require.Equal(t, uint32(0), res.Code)

		msg := shmsg.Message{}
		err := proto.Unmarshal(res.Value, &msg)
		require.Nil(t, err)
		batchConfigMsg := msg.GetBatchConfig()
		require.NotNil(t, batchConfigMsg)
		require.Equal(t, batchConfigMsg.ConfigIndex, testCase.config.ConfigIndex)
		require.Equal(t, batchConfigMsg.StartBatchIndex, testCase.config.StartBatchIndex)
		require.Equal(t, batchConfigMsg.Threshold, testCase.config.Threshold)
		require.Equal(t, len(batchConfigMsg.Keypers), len(testCase.config.Keypers))
		for i := range batchConfigMsg.Keypers {
			address := common.BytesToAddress(batchConfigMsg.Keypers[i])
			require.Equal(t, address, testCase.config.Keypers[i])
		}
	}
}

func TestQueryCheckedIn(t *testing.T) {
	app := NewShutterApp()

	req := abcitypes.RequestQuery{
		Path: "/checkedIn?address=asdf",
	}
	res := app.Query(req)
	require.Equal(t, uint32(1), res.Code)

	address := common.BigToAddress(big.NewInt(0))
	req = abcitypes.RequestQuery{
		Path: fmt.Sprintf("/checkedIn?address=%s", address.Hex()),
	}
	res = app.Query(req)
	require.Equal(t, uint32(0), res.Code)
	require.Equal(t, []byte{0}, res.Value)

	app.Identities[address] = ValidatorPubkey{Ed25519pubkey: ""}
	req = abcitypes.RequestQuery{
		Path: fmt.Sprintf("/checkedIn?address=%s", address.Hex()),
	}
	res = app.Query(req)
	require.Equal(t, uint32(0), res.Code)
	require.Equal(t, []byte{1}, res.Value)
}

func TestCheckInQueryResponseValueParsing(t *testing.T) {
	b, err := ParseCheckInQueryResponseValue([]byte{})
	require.NotNil(t, err)

	b, err = ParseCheckInQueryResponseValue([]byte{0})
	require.Nil(t, err)
	require.False(t, b)

	b, err = ParseCheckInQueryResponseValue([]byte{1})
	require.Nil(t, err)
	require.True(t, b)

	b, err = ParseCheckInQueryResponseValue([]byte{2})
	require.NotNil(t, err)
}

func TestQueryVote(t *testing.T) {
	app := NewShutterApp()

	req := abcitypes.RequestQuery{
		Path: "/vote",
	}
	res := app.Query(req)
	require.Equal(t, uint32(1), res.Code)

	req = abcitypes.RequestQuery{
		Path: "/vote?address=asdf",
	}
	res = app.Query(req)
	require.Equal(t, uint32(1), res.Code)

	address := common.BigToAddress(big.NewInt(0))
	req = abcitypes.RequestQuery{
		Path: fmt.Sprintf("/vote?address=%s", address.Hex()),
	}
	res = app.Query(req)
	require.Equal(t, uint32(0), res.Code)
	require.Equal(t, res.Value, []byte{})

	config := BatchConfig{
		StartBatchIndex: 100,
	}
	app.Voting.AddVote(address, config)
	req = abcitypes.RequestQuery{
		Path: fmt.Sprintf("/vote?address=%s", address.Hex()),
	}
	res = app.Query(req)
	require.Equal(t, uint32(0), res.Code)
	msg := shmsg.Message{}
	err := proto.Unmarshal(res.Value, &msg)
	require.Nil(t, err)
	batchConfigMsg := msg.GetBatchConfig()
	require.NotNil(t, batchConfigMsg)
	require.Equal(t, batchConfigMsg.StartBatchIndex, config.StartBatchIndex)
}
