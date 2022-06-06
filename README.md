# Builder API

Builder API implementing [builder spec](https://github.com/ethereum/builder-specs), with the API separated from the geth codebase in the [original flashbots repository](https://github.com/flashbots/boost-geth-builder). 

Run on your favorite network, including Kiln and local devnet.

Requires forkchoice update to be sent for block building, on public testnets run beacon node modified to send forkchoice update on every slot [example modified beacon client (lighthouse)](https://github.com/flashbots/lighthouse)

Test with [mev-boost](https://github.com/flashbots/mev-boost) and [mev-boost test cli](https://github.com/flashbots/mev-boost/tree/main/cmd/test-cli).

Provides summary page at the listening address' root (http://localhost:28545 by default).

## How to Run

To run the builder, make sure to have go installed:

```bash
make
./build/bin/builder
```

and a relay server will run by default on http://localhost:28545.

[Nethermind](https://github.com/NethermindEth/nethermind) has integrated with this builder API to be a block builder. To produce blocks with `mev-boost` you will need:

* Synced execution client
* Synced beacon client
* Validator
* [mev-boost](https://github.com/flashbots/mev-boost)

Instructions to set up Nethermind with a beacon client can be found [here](https://docs.nethermind.io/nethermind/first-steps-with-nethermind/running-nethermind-post-merge)

To allow Nethermind to be a block builder, make sure the Merge plugin is enabled with the relay url specified in the Nethermind config file:

```
"Merge": {
  "Enabled": true,
  "BuilderRelayUrl": "http://localhost:28545"
}
```

To connect with the kiln network, edit `kiln.cfg` with the relay url and run:

```
dotnet run -c Release --config kiln
```

## How it works

API server has two endpoints that can be called by a compatible execution client (EL):
* `/eth/v1/relay/payload_attributes` - on forkchoice update, EL sends payload attributes to this endpoint. The API changes the payload attributes feeRecipient to the one registered for next slot's validator and returns to the EL.
* `/eth/v1/relay/submit_block` - on new sealed block, EL builds and submits a block with the same fee recipient and gas limit specified in the payload attributes returned previously. Builder will consume the block as the next slot's proposed payload

## Limitations

* Blocks are only built on forkchoice update call from beacon node
* Only works post-Bellatrix, fork version is static
* Does not accept external blocks
* Does not have payload cache, only the latest block is available

## Usage

Builder API options:
```
$ geth --help
BUILDER API OPTIONS:
  --builder.validator_checks               Enable the validator checks
  --builder.secret_key value               Builder API key used for signing headers (default: "0x2fc12ae741f29701f8e30f5de6350766c020cb80768a0ff01e6838ffd2431e11") [$BUILDER_SECRET_KEY]
  --builder.listen_addr value              Listening address for builder endpoint (default: ":28545") [$BUILDER_LISTEN_ADDR]
  --builder.genesis_fork_version value     Gensis fork version (default: "0x02000000") [$BUILDER_GENESIS_FORK_VERSION]
  --builder.bellatrix_fork_version value   Bellatrix fork version (default: "0x02000000") [$BUILDER_BELLATRIX_FORK_VERSION]
  --builder.genesis_validators_root value  Genesis validators root of the network (static). For kiln use 0x99b09fcd43e5905236c370f184056bec6e6638cfc31a323b304fc4aa789cb4ad (default: "0x0000000000000000000000000000000000000000000000000000000000000000") [$BUILDER_GENESIS_VALIDATORS_ROOT]
  --builder.beacon_endpoint value          Beacon endpoint to connect to for beacon chain data (default: "http://127.0.0.1:5052") [$BUILDER_BEACON_ENDPOINT]
```
