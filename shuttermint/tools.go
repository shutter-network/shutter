// +build tools

// Package tools is used to declare and track tool dependencies.  See
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module for
// further information.
package tools

import (
	_ "github.com/ethereum/go-ethereum/cmd/abigen"
	_ "golang.org/x/tools/cmd/stringer"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
