package main

import (
	"fmt"
	"subscription-biling-system/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)
	log := config.NewLogger(viperConfig)
	go config.Whatsmeow()
	config.NewGorm(viperConfig, log)
	config.NewValidator(viperConfig)
	config.NewMidtrans(viperConfig)
	config.NewEmail(viperConfig)

	webPort := viperConfig.GetInt("fiber.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
