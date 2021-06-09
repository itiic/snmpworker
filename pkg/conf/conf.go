package conf

import (
	"github.com/spf13/viper"
)

// Basic config struct
type Config struct {
	Community string `mapstructure:"community"`
	Retry     int    `mapstructure:"retry"`
	Timeout   int    `mapstructure:"timeout"`
	Worker    int    `mapstructure:"worker"`
}

// NewConfig function
func NewConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
