package config

import (
	"github.com/spf13/viper"
)

func NewEmail(viper *viper.Viper) *SMTPConfig {
	sender := viper.GetString("email.sender")
	password := viper.GetString("email.password")
	host := viper.GetString("email.host")
	port := viper.GetString("email.port")

	return &SMTPConfig{
		Host:     host,
		Port:     port,
		Email:    sender,
		Password: password,
	}
}

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}
