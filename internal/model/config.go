package model

import (
	"strings"

	"github.com/ividernvi/algohub-forum/internal/model/options"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	*options.Options
}

func NewConfigFromOptions(opts *options.Options) *Config {
	return &Config{opts}
}

func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/iam")
		viper.AddConfigPath("./conf")
		viper.SetConfigName(defaultName)
	}

	// Use config file from the flag.
	viper.SetConfigType("yaml") // set the type of the configuration to yaml.
	viper.AutomaticEnv()        // read in environment variables that match.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
