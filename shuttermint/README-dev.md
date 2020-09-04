# shuttermint
## Installation

Make sure you have at least go version 1.14 installed. Make sure `PATH`
contains `$GOPATH/bin`. If you didn't set `GOPATH`, it defaults to
`${HOME}/go`.

Run `make` or `make build` to build the executables. The executables
are build in the `bin` directory.

Run `make install-tools` to install additional tools for linting and
compiling the protocol buffer files. The `install-tools` target will
also install the tendermint executable itself. It's currently needed
for generating an initial config file.

## Tests
Run `make test` to run the tests

## Linting
Run `make lint` to run `golangci-lint`.

## Running
Run `make run` to start shuttermint.

