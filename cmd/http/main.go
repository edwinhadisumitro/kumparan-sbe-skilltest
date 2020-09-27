package main

import (
	"kumparan-sbe-skilltest/config"
	"kumparan-sbe-skilltest/helper"
	"kumparan-sbe-skilltest/internal/app"
	"kumparan-sbe-skilltest/middleware"
	"log"
	"os"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/subosito/gotenv"
)

// GitCommit : Current commit hash of application
var GitCommit string

func main() {
	// Load ENV Config
	gotenv.Load(*helper.ProjectFolder + ".env")

	// Load Applicaton Config
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Failed to load config, ERROR : ", err.Error())
	}

	e := echo.New()

	// Middleware
	// Echo's
	e.Use(echoMiddleware.Recover())
	if os.Getenv("ECHO_LOGGER") == "ON" {
		e.Use(echoMiddleware.Logger())
	}

	// Custom
	middl := middleware.NewMiddleware(helper.ProjectFolder)
	e.Use(middl.ValidateContentType)

	app.InitHTTPServer(config, e)

	// Start Server
	log.Printf("Kumparan Senior Backend Engineer Skilltest App (build commit hash: %s) running on port: %s\n", GitCommit, config.AppPort)
	if config.AppPort != "" {
		e.HidePort = true
		e.HideBanner = true
	}
	e.Logger.Fatal(e.Start(":" + config.AppPort))
}
