package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	RC  RedisConfig
}

type APIConfig struct {
	Port string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func init() {
	viper.SetDefault("api.port", "7879")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.RC = RedisConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
	}
	return nil
}

func GetAPIConfig() APIConfig {
	return cfg.API
}

func GetRedisConfig() RedisConfig {
	return cfg.RC
}
