package main

import (
	"fmt"
	"os"
	"path/filepath"

	builder "github.com/avalonche/boost-geth-builder/builder"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
)

var (
	app             = NewApp("the builder api command line interface")
	builderApiFlags = []cli.Flag{
		EnableValidatorChecks,
		SecretKey,
		ListenAddr,
		GenesisForkVersion,
		BellatrixForkVersion,
		GenesisValidatorsRoot,
		BeaconEndpoint,
	}
)

// NewApp creates an app with sane defaults.
func NewApp(usage string) *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = filepath.Base(os.Args[0])
	app.Usage = usage
	return app
}

func builderApi(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	bpConfig := &builder.BuilderConfig{
		EnableValidatorChecks: ctx.GlobalIsSet(EnableValidatorChecks.Name),
		SecretKey:             ctx.GlobalString(SecretKey.Name),
		ListenAddr:            ctx.GlobalString(ListenAddr.Name),
		GenesisForkVersion:    ctx.GlobalString(GenesisForkVersion.Name),
		BellatrixForkVersion:  ctx.GlobalString(BellatrixForkVersion.Name),
		GenesisValidatorsRoot: ctx.GlobalString(GenesisValidatorsRoot.Name),
		BeaconEndpoint:        ctx.GlobalString(BeaconEndpoint.Name),
	}

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)

	if err := builder.Start(bpConfig); err != nil {
		return err
	}

	return nil
}

func init() {
	app.Version = "0.1.0"
	app.Action = builderApi
	app.Flags = builderApiFlags
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
