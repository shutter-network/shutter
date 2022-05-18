package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"syscall/js"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

var (
	uint8Array        js.Value
	uint8ClampedArray js.Value
)

func main() {
	uint8Array = js.Global().Get("Uint8Array")
	uint8ClampedArray = js.Global().Get("Uint8ClampedArray")

	js.Global().Set("shcrypto_encrypt", encrypt)

	select {}
}

var encrypt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) != 4 {
		return encodeResult(nil, fmt.Errorf("expected 4 arguments, got %d", len(args)))
	}
	messageArg := args[0]
	eonPublicKeyArg := args[1]
	epochIDArg := args[2]
	sigmaArg := args[3]

	message, err := decodeMessageArg(messageArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	eonPublicKey, err := decodeEonPublicKeyArg(eonPublicKeyArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	epochID, err := decodeEpochIDArg(epochIDArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	sigma, err := decodeSigmaArg(sigmaArg)
	if err != nil {
		return encodeResult(nil, err)
	}

	encryptedMessage := shcrypto.Encrypt(message, eonPublicKey, epochID, sigma)
	return encodeResult(encryptedMessage.Marshal(), nil)
})

func encodeResult(encryptedMessage []byte, err error) string {
	if err != nil {
		return "Error: " + err.Error()
	}
	return "0x" + hex.EncodeToString(encryptedMessage)
}

func decodeMessageArg(arg js.Value) ([]byte, error) {
	return decodeBytesArg(arg, "message")
}

func decodeEonPublicKeyArg(arg js.Value) (*shcrypto.EonPublicKey, error) {
	b, err := decodeBytesArg(arg, "eonPublicKey")
	if err != nil {
		return nil, err
	}

	p := new(shcrypto.EonPublicKey)
	err = p.Unmarshal(b)
	if err != nil {
		return nil, fmt.Errorf("invalid eon public key: %s", err)
	}

	return p, nil
}

func decodeEpochIDArg(arg js.Value) (*shcrypto.EpochID, error) {
	b, err := decodeBytesArg(arg, "epochID")
	if err != nil {
		return nil, err
	}
	if len(b) != 8 {
		return nil, fmt.Errorf("epochID must be 8 bytes, got %d", len(b))
	}

	i := binary.BigEndian.Uint64(b)
	p := shcrypto.ComputeEpochID(i)
	return p, nil
}

func decodeSigmaArg(arg js.Value) (shcrypto.Block, error) {
	var s shcrypto.Block

	b, err := decodeBytesArg(arg, "sigma")
	if err != nil {
		return s, err
	}
	if len(b) != shcrypto.BlockSize {
		return s, fmt.Errorf("sigma must be %d bytes, got %d", shcrypto.BlockSize, len(b))
	}

	copy(s[:], b)
	return s, nil
}

func decodeBytesArg(arg js.Value, name string) ([]byte, error) {
	if !(arg.InstanceOf(uint8Array) || arg.InstanceOf(uint8ClampedArray)) {
		return nil, fmt.Errorf("argument %s must be of type Uint8Array", name)
	}
	b := make([]byte, arg.Get("length").Int())
	js.CopyBytesToGo(b, arg)
	return b, nil
}
