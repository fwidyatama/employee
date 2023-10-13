package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST" default:"localhost"`
	DBName     string `mapstructure:"DB_NAME" default:"postgres"`
	DBPort     string `mapstructure:"DB_PORT" default:"5432"`
	DBPassword string `mapstructure:"DB_PASSWORD" default:"postgres"`
	DBUser     string `mapstructure:"DB_USER" default:"postgres"`
}

func NewConfig() *Config {

	cfg := &Config{}

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		log.Fatal("failed load env file")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
