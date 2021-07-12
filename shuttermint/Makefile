.POSIX:

PREFIX ?= /usr/local
GO ?= go
GOFLAGS ?=
SHROOT ?= $(CURDIR)
TESTROOT=$(SHROOT)/testchain
BINDIR ?= ./bin

OPENZEPPELIN_CONTRACTS=openzeppelin
OPENZEPPELIN_CONTRACTS_DIR=../contracts/openzeppelin

build:
	@VERSION=`git describe --tags --always --abbrev=4 --dirty`; \
	echo "Building shuttermint $${VERSION}"; \
	${GO} build ${GOFLAGS} -o ${BINDIR} -ldflags "-X github.com/shutter-network/shutter/shuttermint/cmd/shversion.version=$${VERSION}" . ./sandbox/testclient

wasm:
	GOARCH=wasm GOOS=js ${GO} build ${GOFLAGS} -o ${BINDIR}/shcrypto.wasm ./shcryptowasm/main_wasm.go

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

${TESTROOT}:
	${BINDIR}/shuttermint init --dev --root ${TESTROOT}

init: ${TESTROOT}

run: build init
	${BINDIR}/shuttermint chain --config ${TESTROOT}/config/config.toml

test:
	${GO} test ${GOFLAGS} ./...

generate:
	${GO} generate ./...

coverage:
	${GO} test ${GOFLAGS} -covermode=count -coverprofile=coverage.out ./...
	${GO} tool cover -html=coverage.out

clean:
	rm -rf ${TESTROOT}
	rm -f ${BINDIR}/shuttermint ${BINDIR}/testclient ${BINDIR}/prepare

install:
	install -o root -m 0555 ${BINDIR}/shuttermint ${PREFIX}/bin/shuttermint

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
	base=`git merge-base HEAD origin/main`; \
	golangci-lint run --new-from-rev $${base}

abigen:
	solc --allow-paths . ${OPENZEPPELIN_CONTRACTS}=${OPENZEPPELIN_CONTRACTS_DIR} --combined-json=bin,bin-runtime,ast,metadata,abi,srcmap,srcmap-runtime,storage-layout --optimize ../contracts/contracts/binding.sol > contract/combined.json
	abigen --pkg contract --out contract/binding.go --combined-json contract/combined.json

deploy-ganache:
	${GO} run ./sandbox/deploy/deploy.go deploy

.PHONY: build install run init clean test generate install-abigen install-geth install-protoc-gen-go install-golangci-lint install-cobra install-gofumpt install-tools lint lint-changes abigen coverage
