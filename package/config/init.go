package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./package/config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.SetConfigName("private-config")
	viper.SetConfigType("env")
	viper.AddConfigPath(viper.GetString("private_config.path"))
	err = viper.MergeInConfig()
	return err
}

func InitForTests() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../../package/config/")
	fmt.Println()
	err := viper.ReadInConfig()
	return err
}
