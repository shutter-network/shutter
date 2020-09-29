package app

import (
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
	}
	return makeQueryErrorResponse("unknown method")
}

func (app *ShutterApp) queryBatchConfig(vs url.Values) abcitypes.ResponseQuery {
	batchIndexStr := vs.Get("batchIndex")
	if batchIndexStr == "" {
		return makeQueryErrorResponse("missing batch index parameter")
	}

	batchIndex, err := strconv.Atoi(batchIndexStr)
	if err != nil || batchIndex < 0 {
		return makeQueryErrorResponse("batch index not valid integer")
	}

	config := app.getConfig(uint64(batchIndex))
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
