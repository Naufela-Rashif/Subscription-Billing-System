package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type MidtransConfig struct {
	ServerKey   string `mapstructure:"server_key"`
	Environment string `mapstructure:"environment"`
}

func NewMidtrans(v *viper.Viper) (*MidtransConfig, error) {
	var cfg MidtransConfig
	if err := v.Sub("midtrans").Unmarshal(&cfg); err != nil {
		return nil, err
	}
	midtrans.ServerKey = cfg.ServerKey
	if strings.ToLower(cfg.Environment) == "production" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}

	log.Printf("Midtrans Config: ServerKey=%s, Environment=%v", cfg.ServerKey, midtrans.Environment)
	return &cfg, nil
}
