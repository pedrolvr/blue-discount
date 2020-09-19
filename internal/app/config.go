package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Discount DiscountConfig
	DB       DBConfig `mapstructure:"db"`
}

type DiscountConfig struct {
	MaxApplied int32 `mapstructure:"max_applied"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int
	Name     string
	User     string
	Password string
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
