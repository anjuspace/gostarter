package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anjuspace/gostarter/internal/platform/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// configuration holds any kind of configuration that comes from the outside world and
// is necessary for running the application.
type configuration struct {
	// Meaningful values are recommended (eg. production, development, staging, release/123, etc)
	Environment string

	// Turns on some debug functionality
	Debug bool

	// Timeout for graceful shutdown
	ShutdownTimeout time.Duration

	// Log configuration
	Log log.Config
}

// Validate validates the configuration.
func (c configuration) Validate() error {
	if c.Environment == "" {
		return errors.New("environment is required")
	}

	return nil
}

// configure configures some defaults in the Viper instance.
func configure(v *viper.Viper, p *pflag.FlagSet) {
	// Viper settings
	v.AddConfigPath(".")
	v.AddConfigPath(fmt.Sprintf("$%s_CONFIG_DIR/", strings.ToUpper(envPrefix)))

	// Environment variable settings
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Application constants
	v.Set("appName", appName)

	// Global configuration
	v.SetDefault("environment", "production")
	v.SetDefault("debug", false)
	v.SetDefault("shutdownTimeout", 15*time.Second)
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		v.SetDefault("no_color", true)
	}

	// Log configuration
	v.SetDefault("log.format", "json")
	v.SetDefault("log.level", "info")
	v.RegisterAlias("log.noColor", "no_color")
}
