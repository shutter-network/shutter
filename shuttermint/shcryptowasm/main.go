package main

import "syscall/js"

func main() {
	registerCallbacks()

	c := make(chan struct{}, 0)
	<-c
}

func registerCallbacks() {
	shcrypto := make(map[string]interface{})
	shcrypto["mul2"] = mul2

	js.Global().Set("shcrypto", shcrypto)
}

var mul2 = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "invalid number of arguments"
	}
	if args[0].Type() != js.TypeNumber {
		return "invalid argument type"
	}
	i := args[0].Int()

	return 2 * i
})
