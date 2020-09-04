export TMHOME=$(CURDIR)/tmhome

VERSION=$(shell git describe --tags --always)

build:
	go build -o bin -ldflags "-X main.version=${VERSION}" . ./cmd/multik ./cmd/keyper ./cmd/testclient

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

tmhome:
	tendermint init

init: tmhome

run: build init
	./bin/shuttermint --config ${TMHOME}/config/config.toml

test:
	go test -v ./...

clean:
	rm -rf ./tmhome ./data
	rm -f ./shuttermint bin/shuttermint bin/multik bin/testclient bin/keyper

install-tools: install-abigen install-protoc-gen-go install-golangci-lint

install-abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

install-protoc-gen-go:
	go get google.golang.org/protobuf/cmd/protoc-gen-go

install-golangci-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build run init clean test install-abigen install-protoc-gen-go install-golangci-lint install-tools
