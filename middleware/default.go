package middleware

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUri string `mapstructure:"MYSQLDB_LOCAL_URI"`
	RedisUri string `mapstructure:"REDIS_URI"`
	Port string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}