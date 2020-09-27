package main

import (
	"kumparan-sbe-skilltest/config"
	"kumparan-sbe-skilltest/helper"
	"kumparan-sbe-skilltest/internal/app"
	"log"

	"github.com/subosito/gotenv"
)

func main() {
	// Load ENV Config
	gotenv.Load(*helper.ProjectFolder + ".env")

	// Load Applicaton Config
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Failed to load config, ERROR : ", err.Error())
	}

	app.InitNSQSubscriber(config)
}
