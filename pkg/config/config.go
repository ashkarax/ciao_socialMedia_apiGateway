package config_apigw

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	AuthSvcUrl         string `mapstructure:"AUTH_SVC_URL"`
	ApiKey             string `mapstructure:"API_KEY"`
	PostNrelSvcUrl     string `mapstructure:"POSTNREL_SVC_URL"`
	ChatSvcUrl         string `mapstructure:"CHAT_SVC_URL"`
	KafkaPort          string `mapstructure:"KAFKA_PORT"`
	KafkaTopicOneToOne string `mapstructure:"KAFKA_TOPIC_1"`
}

func LoadConfig() (*Config, error) {
	var c Config

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("-------", err)

		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	return &c, nil

}
