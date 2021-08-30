module github.com/shutter-network/shutter/shuttermint

go 1.15

require (
	github.com/DataDog/zstd v1.4.8 // indirect
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/ethereum/go-ethereum v1.10.3
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/karalabe/usb v0.0.0-20191104083709-911d15fe12a9 // indirect
	github.com/kr/pretty v0.2.1
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.4 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.0
	google.golang.org/genproto v0.0.0-20210302174412-5ede27ff9881 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/ini.v1 v1.62.0 // indirect
	gotest.tools/v3 v3.0.3
)

require github.com/shutter-network/shutter/shlib v0.0.0

replace github.com/shutter-network/shutter/shlib => ../shlib
