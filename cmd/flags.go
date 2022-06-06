package main

import (
	"gopkg.in/urfave/cli.v1"
)

// These are all the command line flags supported.
var (
	// Builder API settings
	EnableValidatorChecks = cli.BoolFlag{
		Name:  "validator_checks",
		Usage: "Enable the validator checks",
	}
	SecretKey = cli.StringFlag{
		Name:   "secret_key",
		Usage:  "Builder API key used for signing headers",
		EnvVar: "BUILDER_SECRET_KEY",
		Value:  "0x2fc12ae741f29701f8e30f5de6350766c020cb80768a0ff01e6838ffd2431e11",
	}
	ListenAddr = cli.StringFlag{
		Name:   "listen_addr",
		Usage:  "Listening address for builder endpoint",
		EnvVar: "BUILDER_LISTEN_ADDR",
		Value:  ":28545",
	}
	GenesisForkVersion = cli.StringFlag{
		Name:   "genesis_fork_version",
		Usage:  "Gensis fork version. For kiln use 0x70000069",
		EnvVar: "BUILDER_GENESIS_FORK_VERSION",
		Value:  "0x00000000",
	}
	BellatrixForkVersion = cli.StringFlag{
		Name:   "bellatrix_fork_version",
		Usage:  "Bellatrix fork version. For kiln use 0x70000071",
		EnvVar: "BUILDER_BELLATRIX_FORK_VERSION",
		Value:  "0x02000000",
	}
	GenesisValidatorsRoot = cli.StringFlag{
		Name:   "genesis_validators_root",
		Usage:  "Genesis validators root of the network. For kiln use 0x99b09fcd43e5905236c370f184056bec6e6638cfc31a323b304fc4aa789cb4ad",
		EnvVar: "BUILDER_GENESIS_VALIDATORS_ROOT",
		Value:  "0x0000000000000000000000000000000000000000000000000000000000000000",
	}
	BeaconEndpoint = cli.StringFlag{
		Name:   "beacon_endpoint",
		Usage:  "Beacon endpoint to connect to for beacon chain data",
		EnvVar: "BUILDER_BEACON_ENDPOINT",
		Value:  "http://127.0.0.1:5052",
	}
)
