package cmd

import (
	"context"
	"io/ioutil"

	"github.com/pastelnetwork/go-commons/cli"
	"github.com/pastelnetwork/go-commons/configurer"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/go-commons/log"
	"github.com/pastelnetwork/go-commons/log/hooks"
	"github.com/pastelnetwork/go-commons/sys"
	"github.com/pastelnetwork/go-commons/version"
	"github.com/pastelnetwork/go-pastel"
	"github.com/pastelnetwork/walletnode/api"
	"github.com/pastelnetwork/walletnode/configs"
	"github.com/pastelnetwork/walletnode/services/artwork/register"
	"github.com/pastelnetwork/walletnode/storage/memory"
)

const (
	appName  = "walletnode"
	appUsage = "WalletNode" // TODO: Write a clear description.

	defaultConfigFile = ""
)

// NewApp inits a new command line interface.
func NewApp() *cli.App {
	configFile := defaultConfigFile
	config := configs.New()

	app := cli.NewApp(appName)
	app.SetUsage(appUsage)
	app.SetVersion(version.Version())

	app.AddFlags(
		// Main
		cli.NewFlag("config-file", &configFile).SetUsage("Set `path` to the config file.").SetValue(configFile).SetAliases("c"),
		cli.NewFlag("log-level", &config.LogLevel).SetUsage("Set the log `level`.").SetValue(config.LogLevel),
		cli.NewFlag("log-file", &config.LogFile).SetUsage("The log `file` to write to."),
		cli.NewFlag("quiet", &config.Quiet).SetUsage("Disallows log output to stdout.").SetAliases("q"),
		// Rest
		cli.NewFlag("swagger", &config.Rest.Swagger).SetUsage("Enable Swagger UI."),
	)

	app.SetActionFunc(func(args []string) error {
		ctx := context.TODO()

		if configFile != "" {
			if err := configurer.ParseFile(configFile, config); err != nil {
				return err
			}
		}

		if config.Quiet {
			log.SetOutput(ioutil.Discard)
		} else {
			log.SetOutput(app.Writer)
		}

		if config.LogFile != "" {
			fileHook := hooks.NewFileHook(config.LogFile)
			log.AddHook(fileHook)
		}

		if err := log.SetLevelName(config.LogLevel); err != nil {
			return errors.Errorf("--log-level %q, %s", config.LogLevel, err)
		}

		return action(ctx, config)
	})

	return app
}

func action(ctx context.Context, config *configs.Config) error {
	log.Debug("[app] start")
	defer log.Debug("[app] end")

	log.Debugf("[app] config: %s", config)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sys.RegisterInterruptHandler(cancel, func() {
		log.Info("[app] Interrupt signal received. Gracefully shutting down...")
	})

	// periphery
	pastel := pastel.NewClient(config.Pastel)
	db := memory.NewKeyValue()

	// services
	artwork := register.NewService(db, pastel)

	// api
	api := api.New(config.Rest,
		api.NewArtwork(artwork),
		api.NewSwagger(),
	)

	return run(ctx, artwork, api)
}
