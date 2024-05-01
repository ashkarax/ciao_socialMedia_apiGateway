package config_apigw

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL"`
	ApiKey     string `mapstructure:"API_KEY"`

	//ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	//OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (*Config, error) {
	var c Config

	viper.AddConfigPath("./pkg/config/envs")
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
