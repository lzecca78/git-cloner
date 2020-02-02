package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *GitConfig

type GitConfig struct {
	Repos []Repo `mapstructure:"repos"`
}

type Repo struct {
	LocalDir   string `mapstructure:"local_dir"`
	Git_Remote string `mapstructure:"git_remote"`
}

func GetConfig(cfgFile string) *GitConfig {
	// Reads the configuration file
	gitConfig, err := InitCfg(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	return gitConfig
}

func InitCfg(configFile string) (*GitConfig, error) {

	viper.SetConfigFile(configFile)

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error reading config using config file: %v", viper.ConfigFileUsed())
	}
	// Populate the conf with configuration data
	conf := new(GitConfig)
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return conf, nil
}
