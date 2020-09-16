package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Discount DiscountConfig
}

type DiscountConfig struct {
	MaxApplied int32 `mapstructure:"max_applied"`
}

func ReadConfig(name, path string) (Config, error) {
	var c Config

	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return c, fmt.Errorf("error config file: %w \n", err)
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return c, fmt.Errorf("unable to decode config: %w \n", err)
	}

	return c, nil
}
