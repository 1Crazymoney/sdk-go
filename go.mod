module github.com/InjectiveLabs/sdk-go

go 1.15

require (
	github.com/InjectiveLabs/injective-core v0.0.0-20210113231547-583470c27cc5
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/bugsnag/panicwrap v1.2.0 // indirect
	github.com/cosmos/cosmos-sdk v0.40.0
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/karalabe/usb v0.0.0-20191104083709-911d15fe12a9 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/regen-network/cosmos-proto v0.3.0
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.1
	github.com/tyler-smith/go-bip39 v1.0.2
	github.com/xlab/suplog v1.0.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
