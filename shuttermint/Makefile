export TMHOME=$(CURDIR)/tmhome

build:
	go build

protoc:
	protoc shmsg/shmsg.proto --go_out=shmsg/

tmhome:
	tendermint init

init: tmhome

run: build init
	./shuttermint --config ${TMHOME}/config/config.toml

test:
	go test -v ./...

clean:
	rm -rf ./shuttermint ./tmhome

install-abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

install-protoc-gen-go:
	go get google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: build run init clean test install-abigen
