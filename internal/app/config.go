package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Discount DiscountConfig
	Service  ServiceConfig
	DB       DBConfig
}

type ServiceConfig struct {
	Port int32
}

type DiscountConfig struct {
	MaxApplied int32 `mapstructure:"max_applied"`
}

type DBConfig struct {
	Host     string
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
		return c, fmt.Errorf("read config: %w \n", err)
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return c, fmt.Errorf("unable to decode config: %w \n", err)
	}

	return c, nil
}
