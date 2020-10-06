TESTROOT=$(CURDIR)/testchain

VERSION=$(shell git describe --tags --always --abbrev=4 --dirty)

build:
	go build -o bin -ldflags "-X main.version=${VERSION} -X github.com/brainbot-com/shutter/shuttermint/cmd.version=${VERSION}" . ./sandbox/testclient ./sandbox/deploy

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

${TESTROOT}:
	bin/shuttermint init --root ${TESTROOT}

init: ${TESTROOT}

run: build init
	./bin/shuttermint run --config ${TESTROOT}/config/config.toml

test:
	go test ./...

coverage:
	go test -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	rm -rf ${TESTROOT}
	rm -f ./shuttermint bin/shuttermint bin/multik bin/testclient bin/keyper bin/deploy

install-tools: install-abigen install-protoc-gen-go install-golangci-lint install-cobra install-gofumports

install-geth:
	go install github.com/ethereum/go-ethereum/cmd/geth

install-abigen:
	go install github.com/ethereum/go-ethereum/cmd/abigen

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

install-cobra:
	go install github.com/spf13/cobra/cobra

install-gofumports:
	go install mvdan.cc/gofumpt/gofumports

lint:
	golangci-lint run --tests

lint-changes:
	golangci-lint run --new-from-rev origin/master

abigen:
	abigen --pkg contract --out contract/binding.go --sol ../contracts/contracts/ConfigContract.sol --sol ../contracts/contracts/KeyBroadcastContract.sol --sol ../contracts/contracts/BatcherContract.sol

deploy-ganache:
	go run ./sandbox/deploy/deploy.go deploy

.PHONY: build run init clean test install-abigen install-geth install-protoc-gen-go install-golangci-lint install-cobra install-gofumports install-tools lint lint-changes abigen coverage
