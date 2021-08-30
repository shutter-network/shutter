package main

import (
	"syscall/js"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/shutter-network/shutter/shlib/shcrypto"
)

func main() {
	registerCallbacks()

	c := make(chan struct{}, 0)
	<-c
}

func registerCallbacks() {
	shcrypto := make(map[string]interface{})

	shcrypto["encrypt"] = encrypt

	js.Global().Set("shcrypto", shcrypto)
}

var encrypt = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	returnValue := func(encoded []byte, err error) map[string]interface{} {
		m := make(map[string]interface{})
		m["encryptedMessage"] = hexutil.Encode(encoded)
		if err != nil {
			m["error"] = err.Error()
		} else {
			m["error"] = nil
		}
		return m
	}
	errorReturnValue := func(err error) map[string]interface{} {
		return returnValue([]byte{}, err)
	}

	if len(args) != 4 {
		return errorReturnValue(errors.Errorf("expected 4 arguments, got %d", len(args)))
	}
	messageJS := args[0]
	eonPublicKeyJS := args[1]
	epochIndexJS := args[2]
	sigmaJS := args[3]

	if err := validateMessage(messageJS); err != nil {
		return errorReturnValue(err)
	}
	if err := validateEonPublicKey(eonPublicKeyJS); err != nil {
		return errorReturnValue(err)
	}
	if err := validateEpochIndex(epochIndexJS); err != nil {
		return errorReturnValue(err)
	}
	if err := validateSigma(sigmaJS); err != nil {
		return errorReturnValue(err)
	}

	message := make([]byte, messageJS.Length())
	js.CopyBytesToGo(message, messageJS)

	eonPublicKeyBytes := make([]byte, eonPublicKeyJS.Length())
	js.CopyBytesToGo(eonPublicKeyBytes, eonPublicKeyJS)
	eonPublicKey := new(shcrypto.EonPublicKey)
	err := eonPublicKey.Unmarshal(eonPublicKeyBytes)
	if err != nil {
		return errorReturnValue(errors.Wrap(err, "failed to decode eon public key"))
	}

	epochIndex := uint64(epochIndexJS.Int())
	epochID := shcrypto.ComputeEpochID(epochIndex)

	var sigma shcrypto.Block
	js.CopyBytesToGo(sigma[:], sigmaJS)

	m := shcrypto.Encrypt(message, eonPublicKey, epochID, sigma)
	encoded := m.Marshal()
	return returnValue(encoded, nil)
})

func validateMessage(v js.Value) error {
	return validateUint8Array(v)
}

func validateEonPublicKey(v js.Value) error {
	return validateUint8Array(v)
}

func validateEpochIndex(v js.Value) error {
	if v.Type() != js.TypeNumber {
		return errors.Errorf("expected number, got non-number")
	}
	if v.Int() < 0 {
		return errors.Errorf("epoch index must not be negative, got %d", v.Int())
	}
	return nil
}

func validateSigma(v js.Value) error {
	return validateBlock(v)
}

func validateBlock(v js.Value) error {
	if err := validateUint8Array(v); err != nil {
		return err
	}
	if v.Length() != 32 {
		return errors.Errorf("expected array of length 32, got %d", v.Length())
	}
	return nil
}

func validateUint8Array(v js.Value) error {
	if v.Type() != js.TypeObject {
		return errors.Errorf("expected Uint8Array, but value is not an object")
	}
	constructor := v.Get("constructor")
	if constructor.Type() != js.TypeFunction {
		return errors.Errorf("expected Uint8Array, but value constructor is not a function")
	}
	name := constructor.Get("name")
	if name.Type() != js.TypeString {
		return errors.Errorf("expected Uint8Array, but name is not a string")
	}
	if name.String() != "Uint8Array" {
		return errors.Errorf("expected Uint8Array, but name is %s", name)
	}
	return nil
}
