.POSIX:

GO ?= go
SHROOT ?= $(CURDIR)
TESTROOT=$(SHROOT)/testchain

OPENZEPPELIN_CONTRACTS=OpenZeppelin/openzeppelin-contracts@3.3.0
OPENZEPPELIN_CONTRACTS_DIR=../contracts/openzeppelin

build:
	@VERSION=`git describe --tags --always --abbrev=4 --dirty`; \
	echo "Building shuttermint $${VERSION}"; \
	${GO} build -o bin -ldflags "-X github.com/brainbot-com/shutter/shuttermint/cmd/shversion.version=$${VERSION}" . ./sandbox/testclient ./sandbox/prepare

wasm:
	GOARCH=wasm GOOS=js ${GO} build -o bin/shcrypto.wasm ./shcryptowasm/main_wasm.go

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

${TESTROOT}:
	bin/shuttermint init --dev --root ${TESTROOT}

init: ${TESTROOT}

run: build init
	./bin/shuttermint chain --config ${TESTROOT}/config/config.toml

test:
	${GO} test ./...

generate:
	${GO} generate ./...

coverage:
	${GO} test -covermode=count -coverprofile=coverage.out ./...
	${GO} tool cover -html=coverage.out

clean:
	rm -rf ${TESTROOT}
	rm -f bin/shuttermint bin/testclient bin/keyper bin/deploy bin/prepare

install-tools: install-abigen install-protoc-gen-go install-golangci-lint install-cobra install-gofumpt install-stringer install-gci

install-geth:
	${GO} install github.com/ethereum/go-ethereum/cmd/geth@latest

install-abigen:
	${GO} install github.com/ethereum/go-ethereum/cmd/abigen

install-protoc-gen-go:
	${GO} install google.golang.org/protobuf/cmd/protoc-gen-go

install-golangci-lint:
	${GO} install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-cobra:
	${GO} install github.com/spf13/cobra/cobra@latest

install-gofumpt:
	${GO} install mvdan.cc/gofumpt@latest

install-stringer:
	${GO} install golang.org/x/tools/cmd/stringer

install-gci:
	${GO} install github.com/daixiang0/gci@latest

lint:
	golangci-lint run --tests

lint-changes:
	golangci-lint run --new-from-rev origin/master

abigen:
	solc --allow-paths . ${OPENZEPPELIN_CONTRACTS}=${OPENZEPPELIN_CONTRACTS_DIR} --combined-json=bin,bin-runtime,ast,metadata,abi,srcmap,srcmap-runtime,storage-layout --optimize ../contracts/contracts/binding.sol > contract/combined.json
	abigen --pkg contract --out contract/binding.go --combined-json contract/combined.json

deploy-ganache:
	${GO} run ./sandbox/deploy/deploy.go deploy

.PHONY: build run init clean test generate install-abigen install-geth install-protoc-gen-go install-golangci-lint install-cobra install-gofumpt install-tools lint lint-changes abigen coverage
