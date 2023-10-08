package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Logger     Logger     `mapstructure:"logger"      validate:"required"`
	HTTPServer HTTPServer `mapstructure:"http_server" validate:"required"`
	Redis      Redis      `mapstructure:"redis"   validate:"required"`
}

type (
	Logger struct {
		Level *int8 `mapstructure:"level" validate:"required"`
	}

	Redis struct {
		Host     string `mapstructure:"host" validate:"required"`
		DBName   int    `mapstructure:"dbname" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
	}

	HTTPServer struct {
		Address string `mapstructure:"address" validate:"required"`
	}
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	defaultLevel := int8(-1)

	viper.SetDefault("logger.level", &defaultLevel)

	viper.SetDefault("redis.host", "127.0.0.1:6379")
	viper.SetDefault("redis.dbname", 0)
	viper.SetDefault("redis.password", "")
	//
	viper.SetDefault("http_server.address", "0.0.0.0:8080")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
