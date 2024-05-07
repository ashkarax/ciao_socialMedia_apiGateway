package main

import (
	"fmt"
	"log"

	di_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/di"
	responsemodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/response_models"
	apigw_config "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	di_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/di"
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

	app.Use(func(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-Key")
		if apiKey != config.ApiKey {
			return c.Status(fiber.StatusUnauthorized).
				JSON(responsemodels_auth_apigw.CommonResponse{
					StatusCode: fiber.StatusUnauthorized,
					Message:    "request failed",
					Error:      "Invalid API key or No API key found in request",
				})
		}
		return c.Next()
	})

	middleware, err := di_auth_apigw.InitAuthClient(app, config)
	if err != nil {
		log.Fatal(err)
	}

	err = di_postnrel_apigw.InitPostNrelClient(app, config, middleware)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(config.Port)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
