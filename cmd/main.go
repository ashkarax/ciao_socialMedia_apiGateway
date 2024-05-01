package main

import (
	"fmt"
	"log"

	di_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/di"
	apigw_config "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config, err := apigw_config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{TimeFormat: "2006/01/02 15:04:05"}))

	_, err = di_auth_apigw.InitAuthClient(app, config)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(config.Port)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)

	}
}
