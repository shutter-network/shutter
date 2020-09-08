TESTROOT=$(CURDIR)/testchain

VERSION=$(shell git describe --tags --always --abbrev=4 --dirty)

build:
	go build -o bin -ldflags "-X main.version=${VERSION} -X github.com/brainbot-com/shutter/shuttermint/cmd.version=${VERSION}" . ./cmd/multik ./cmd/testclient

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

${TESTROOT}:
	bin/shuttermint init --root ${TESTROOT}

init: ${TESTROOT}

run: build init
	./bin/shuttermint run --config ${TESTROOT}/config/config.toml

test:
	go test -v ./...

clean:
	rm -rf ${TESTROOT}
	rm -f ./shuttermint bin/shuttermint bin/multik bin/testclient bin/keyper

install-tools: install-abigen install-protoc-gen-go install-golangci-lint

install-abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

install-protoc-gen-go:
	go get google.golang.org/protobuf/cmd/protoc-gen-go

install-golangci-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run --tests

abigen:
	abigen --pkg contract --out contract/config.go --sol ../contracts/contracts/ConfigContract.sol

.PHONY: build run init clean test install-abigen install-protoc-gen-go install-golangci-lint install-tendermint install-tools lint abigen
