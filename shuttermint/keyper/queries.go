package keyper

import "fmt"
import tmbytes "github.com/tendermint/tendermint/libs/bytes"
import "github.com/golang/protobuf/proto"
import "github.com/tendermint/tendermint/rpc/client"
import "github.com/brainbot-com/shutter/shuttermint/app"
import "github.com/brainbot-com/shutter/shuttermint/shmsg"

func queryBatchConfig(cl client.Client, batchIndex uint64) (app.BatchConfig, error) {
	path := fmt.Sprintf("/configs?batchIndex=%d", batchIndex)
	data := tmbytes.HexBytes([]byte{})
	res, err := cl.ABCIQuery(path, data)

	if err != nil {
		return app.BatchConfig{}, err
	}
	if res.Response.Code != 0 {
		return app.BatchConfig{}, fmt.Errorf("Query returned error %d: %s", res.Response.Code, res.Response.Log)
	}

	msg := shmsg.Message{}
	err = proto.Unmarshal(res.Response.Value, &msg)
	if err != nil {
		return app.BatchConfig{}, err
	}
	batchConfigMsg := msg.GetBatchConfig()
	if batchConfigMsg == nil {
		return app.BatchConfig{}, fmt.Errorf("Received unexpected message type")
	}

	bc, err := app.BatchConfigFromMessage(batchConfigMsg)
	if err != nil {
		return app.BatchConfig{}, err
	}

	return bc, nil
}
