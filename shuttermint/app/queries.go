package app

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"
)

func makeQueryErrorResponse(msg string) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{
		Code: 1,
		Log:  msg,
	}
}

func isValidRequestURL(u *url.URL) bool {
	return u.Scheme == "" && u.Opaque == "" && u.User.String() == "" && u.Host == "" && u.Fragment == ""
}

func (app *ShutterApp) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	requestURL, err := url.Parse(req.Path)
	if err != nil {
		return makeQueryErrorResponse("invalid request url")
	}
	if !isValidRequestURL(requestURL) {
		return makeQueryErrorResponse("invalid request url")
	}

	switch requestURL.Path {
	case "/configs":
		return app.queryBatchConfig(requestURL.Query())
	case "/checkedIn":
		return app.queryCheckedIn(requestURL.Query())
	case "/vote":
		return app.queryVote(requestURL.Query())
	}
	return makeQueryErrorResponse("unknown method")
}

func (app *ShutterApp) queryBatchConfig(vs url.Values) abcitypes.ResponseQuery {
	batchIndexStr := vs.Get("batchIndex")
	lastConfigStr := vs.Get("last")
	if batchIndexStr == "" && lastConfigStr == "" {
		return makeQueryErrorResponse("missing parameter batchIndex or last")
	}
	if batchIndexStr != "" && lastConfigStr != "" {
		return makeQueryErrorResponse("both batchIndex and last parameter given")
	}

	var config *BatchConfig
	if batchIndexStr != "" {
		batchIndex, err := strconv.Atoi(batchIndexStr)
		if err != nil || batchIndex < 0 {
			return makeQueryErrorResponse("batch index not valid integer")
		}
		config = app.getConfig(uint64(batchIndex))
	} else if lastConfigStr != "" {
		if lastConfigStr != "true" {
			return makeQueryErrorResponse("last parameter must be true if given")
		}
		config = app.Configs[len(app.Configs)-1]
	} else {
		panic("unreachable")
	}

	configMsg := config.Message()
	configBytes, err := proto.Marshal(&configMsg)
	if err != nil {
		return makeQueryErrorResponse("error encoding message")
	}

	return abcitypes.ResponseQuery{
		Code:  0,
		Value: configBytes,
	}
}

func (app *ShutterApp) queryCheckedIn(vs url.Values) abcitypes.ResponseQuery {
	addressStr := vs.Get("address")
	if addressStr == "" {
		return makeQueryErrorResponse("missing address parameter")
	}

	address := common.HexToAddress(addressStr)
	if addressStr != address.Hex() {
		return makeQueryErrorResponse("invalid address")
	}

	var resultByte byte
	if _, ok := app.Identities[address]; ok {
		resultByte = 1
	}

	return abcitypes.ResponseQuery{
		Code:  0,
		Value: []byte{resultByte},
	}
}

func (app *ShutterApp) queryVote(vs url.Values) abcitypes.ResponseQuery {
	addressStr := vs.Get("address")
	if addressStr == "" {
		return makeQueryErrorResponse("missing address parameter")
	}

	address := common.HexToAddress(addressStr)
	if addressStr != address.Hex() {
		return makeQueryErrorResponse("invalid address")
	}

	index, voted := app.ConfigVoting.Votes[address]
	if !voted {
		return abcitypes.ResponseQuery{
			Code:  0,
			Value: []byte{},
		}
	}

	config := app.ConfigVoting.Candidates[index]
	configMsg := config.Message()
	configBytes, err := proto.Marshal(&configMsg)
	if err != nil {
		return makeQueryErrorResponse("error encoding message")
	}

	return abcitypes.ResponseQuery{
		Code:  0,
		Value: configBytes,
	}
}

// ParseCheckInQueryResponseValue interprets the response value of the checkin query
func ParseCheckInQueryResponseValue(v []byte) (bool, error) {
	if len(v) != 1 {
		return false, fmt.Errorf("check in response must be single byte, got %d", len(v))
	}
	switch v[0] {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("check in response byte must be either 0 or 1, got %v", v[0])
	}
}
