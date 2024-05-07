package di_auth_apigw

import (
	"log"

	handler_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/handler"
	client_apigw_auth "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/client"
	router_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/router"
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func InitAuthClient(app *fiber.App, config *config_apigw.Config) (*middleware_auth_apigw.Middleware, error) {

	client, err := client_apigw_auth.InitAuthClient(config)
	if err != nil {
		log.Fatal(err)
	}

	middleware := middleware_auth_apigw.NewAuthMiddleware(client)

	userHandler := handler_auth_apigw.NewAuthUserHandler(client)

	router_auth_apigw.AuthUserRoutes(app, userHandler, middleware)

	return middleware, nil
}
