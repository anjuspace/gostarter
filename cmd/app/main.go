package main

import (
	"fmt"
	"os"

	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/anjuspace/gostarter/internal/platform/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Provisioned by ldflags
var (
	version    string
	commitHash string
	buildDate  string
)

func main() {
	// load application configurations
	v, p := viper.New(), pflag.NewFlagSet(friendlyAppName, pflag.ExitOnError)

	configure(v, p)

	p.String("config", "", "Configuration file")
	p.Bool("version", false, "Show version information")

	_ = p.Parse(os.Args[1:])

	if v, _ := p.GetBool("version"); v {
		fmt.Printf("%s version %s (%s) built on %s\n", friendlyAppName, version, commitHash, buildDate)

		os.Exit(0)
	}

	if c, _ := p.GetString("config"); c != "" {
		v.SetConfigFile(c)
	}

	err := v.ReadInConfig()
	_, configFileNotFound := err.(viper.ConfigFileNotFoundError)
	if !configFileNotFound {
		emperror.Panic(errors.Wrap(err, "failed to read configuration"))
	}

	var config configuration
	err = v.Unmarshal(&config)
	emperror.Panic(errors.Wrap(err, "failed to unmarshal configuration"))
	// Create logger (first thing after configuration loading)
	logger := log.NewLogger(config.Log)

	err = config.Validate()
	if err != nil {
		logger.Error(err.Error())

		os.Exit(3)
	}

	logger.Info("Application started...")
}
