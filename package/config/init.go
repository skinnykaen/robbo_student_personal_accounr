package config

import (
	"github.com/spf13/viper"
)

//type Env struct {
//	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
//	GraphqlServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
//
//	PostgresDsn string `mapstructure:"POSTGRES_DSN"`
//
//	AuthHashSalt                string        `mapstructure:"AUTH_HASH_SALT"`
//	AuthAccessSigningKey        string        `mapstructure:"AUTH_ACCESS_SIGNING_KEY"`
//	AuthRefreshSigningKey       string        `mapstructure:"AUTH_REFRESH_SIGNING_KEY"`
//	AuthAccessTokenTtl          time.Duration `mapstructure:"AUTH_ACCESS_TOKEN_TTL"`
//	AuthRefreshTokenTtl         time.Duration `mapstructure:"AUTH_REFRESH_TOKEN_TTL"`
//	AuthPassResetCodeExpiration time.Duration `mapstructure:"AUTH_PASS_RESET_CODE_EXPIRATION"`
//
//	SmtpServerHost string `mapstructure:"SMTP_SERVER_HOST"`
//	SmtpServerPort string `mapstructure:"SMTP_SERVER_PORT"`
//	Username       string `mapstructure:"MAIL_USERNAME"`
//	Password       string `mapstructure:"MAIL_PASSWORD"`
//
//	ApiToken               string `mapstructure:"API_TOKEN"`
//	ApiTokenId             string `mapstructure:"API_CLIENT_ID"`
//	ApiTokenExpirationTime int64  `mapstructure:"API_TOKEN_EXPIRATION_TIME"`
//	ApiClientSecret        string `mapstructure:"API_CLIENT_SECRET"`
//}

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
	viper.SetConfigName("config-test")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../package/config")

	err := viper.ReadInConfig()
	return err
}
