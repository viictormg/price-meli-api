package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbHost  string   `mapstructure:"DB_HOST"`
	DbPort  string   `mapstructure:"DB_PORT"`
	DbUser  string   `mapstructure:"DB_USER"`
	Brokers []string `mapstructure:"KAFKA_BROKERS"`
	Topic   string   `mapstructure:"KAFKA_TOPIC"`
}

type KafkaConfig struct {
	Brokers []string `mapstructure:"KAFKA_BROKERS"`
	Topic   string   `mapstructure:"KAFKA_TOPIC"`
}

func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	fmt.Println(config.DbHost)

	return &config
}

func (c *Config) GeKafkaConfg() *KafkaConfig {
	return &KafkaConfig{
		Brokers: c.Brokers,
		Topic:   c.Topic,
	}
}
