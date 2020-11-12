package keyper

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/protobuf/proto"

	"github.com/brainbot-com/shutter/shuttermint/app"
	"github.com/brainbot-com/shutter/shuttermint/shmsg"
)

func queryConfig(cl client.Client, path string) (app.BatchConfig, error) {
	data := tmbytes.HexBytes([]byte{})
	res, err := cl.ABCIQuery(context.TODO(), path, data)
	if err != nil {
		return app.BatchConfig{}, err
	}
	if res.Response.Code != 0 {
		return app.BatchConfig{}, fmt.Errorf("query returned error %d: %s", res.Response.Code, res.Response.Log)
	}

	msg := shmsg.Message{}
	err = proto.Unmarshal(res.Response.Value, &msg)
	if err != nil {
		return app.BatchConfig{}, err
	}
	batchConfigMsg := msg.GetBatchConfig()
	if batchConfigMsg == nil {
		return app.BatchConfig{}, fmt.Errorf("received unexpected message type")
	}

	bc, err := app.BatchConfigFromMessage(batchConfigMsg)
	if err != nil {
		return app.BatchConfig{}, err
	}

	return bc, nil
}

func queryBatchConfig(cl client.Client, batchIndex uint64) (app.BatchConfig, error) {
	path := fmt.Sprintf("/configs?batchIndex=%d", batchIndex)
	return queryConfig(cl, path)
}

func queryLastBatchConfig(cl client.Client) (app.BatchConfig, error) {
	path := "/configs?last=true"
	return queryConfig(cl, path)
}

func queryCheckedIn(cl client.Client, address common.Address) (bool, error) {
	path := fmt.Sprintf("/checkedIn?address=%s", address.Hex())
	res, err := cl.ABCIQuery(context.TODO(), path, []byte{})
	if err != nil {
		return false, err
	}
	if res.Response.Code != 0 {
		return false, fmt.Errorf("check in query failed with code %d", res.Response.Code)
	}
	return app.ParseCheckInQueryResponseValue(res.Response.Value)
}

func queryVote(cl client.Client, address common.Address) (app.BatchConfig, bool, error) {
	path := fmt.Sprintf("/vote?address=%s", address.Hex())
	res, err := cl.ABCIQuery(context.TODO(), path, []byte{})
	if err != nil {
		return app.BatchConfig{}, false, err
	}
	if res.Response.Code != 0 {
		return app.BatchConfig{}, false, fmt.Errorf("vote query failed with code %d", res.Response.Code)
	}
	if len(res.Response.Value) == 0 {
		// no vote
		return app.BatchConfig{}, false, nil
	}

	msg := shmsg.Message{}
	err = proto.Unmarshal(res.Response.Value, &msg)
	if err != nil {
		return app.BatchConfig{}, false, err
	}
	batchConfigMsg := msg.GetBatchConfig()
	if batchConfigMsg == nil {
		return app.BatchConfig{}, false, fmt.Errorf("received unexpected message type")
	}
	batchConfig, err := app.BatchConfigFromMessage(batchConfigMsg)
	if err != nil {
		return app.BatchConfig{}, false, err
	}
	return batchConfig, true, nil
}
