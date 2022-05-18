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

	js.Global().Set("shcrypto_decrypt", decrypt)

	select {}
}

var decrypt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return encodeResult(nil, fmt.Errorf("expected 4 arguments, got %d", len(args)))
	}
	encryptedMessageArg := args[0]
	decryptionKeyArg := args[1]

	encryptedMessage, err := decodeEncryptedMessageArg(encryptedMessageArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	decryptionKey, err := decodeDecryptionKeyArg(decryptionKeyArg)
	if err != nil {
		return encodeResult(nil, err)
	}

	message, err := encryptedMessage.Decrypt(decryptionKey)
	if err != nil {
		return encodeResult(nil, fmt.Errorf("failed to decrypt message: %s", err))
	}
	return encodeResult(message, nil)
})

var verifyDecryptionKey = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) != 3 {
		return encodeResult(nil, fmt.Errorf("expected 4 arguments, got %d", len(args)))
	}
	decryptionKeyArg := args[0]
	eonPublicKeyArg := args[1]
	epochIDArg := args[2]

	decryptionKey, err := decodeDecryptionKeyArg(decryptionKeyArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	eonPublicKey, err := decodeEonPublicKeyArg(eonPublicKeyArg)
	if err != nil {
		return encodeResult(nil, err)
	}
	epochID, err := decodeUint64Arg(epochIDArg, "epochID")
	if err != nil {
		return encodeResult(nil, err)
	}

	ok, err := shcrypto.VerifyEpochSecretKey(decryptionKey, eonPublicKey, epochID)
	if err != nil {
		return encodeResult(nil, err)
	}
	if ok {
		return "true"
	}
	return "false"
})

func encodeResult(encryptedMessage []byte, err error) string {
	if err != nil {
		return "Error: " + err.Error()
	}
	return "0x" + hex.EncodeToString(encryptedMessage)
}

func decodeEncryptedMessageArg(arg js.Value) (*shcrypto.EncryptedMessage, error) {
	b, err := decodeBytesArg(arg, "encryptedMessage")
	if err != nil {
		return nil, err
	}

	m := new(shcrypto.EncryptedMessage)
	err = m.Unmarshal(b)
	if err != nil {
		return nil, fmt.Errorf("invalid encrypted message: %s", err)
	}

	return m, nil
}

func decodeDecryptionKeyArg(arg js.Value) (*shcrypto.EpochSecretKey, error) {
	b, err := decodeBytesArg(arg, "eonPublicKey")
	if err != nil {
		return nil, err
	}

	k := new(shcrypto.EpochSecretKey)
	err = k.Unmarshal(b)
	if err != nil {
		return nil, fmt.Errorf("invalid decryption key: %s", err)
	}

	return k, nil
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
	i, err := decodeUint64Arg(arg, "epochID")
	if err != nil {
		return nil, err
	}
	p := shcrypto.ComputeEpochID(i)
	return p, nil
}

func decodeUint64Arg(arg js.Value, name string) (uint64, error) {
	b, err := decodeBytesArg(arg, name)
	if err != nil {
		return 0, err
	}
	if len(b) != 8 {
		return 0, fmt.Errorf("%s must be 8 bytes, got %d", name, len(b))
	}
	return binary.BigEndian.Uint64(b), nil
}

func decodeBytesArg(arg js.Value, name string) ([]byte, error) {
	if !(arg.InstanceOf(uint8Array) || arg.InstanceOf(uint8ClampedArray)) {
		return nil, fmt.Errorf("argument %s must be of type Uint8Array", name)
	}
	b := make([]byte, arg.Get("length").Int())
	js.CopyBytesToGo(b, arg)
	return b, nil
}
