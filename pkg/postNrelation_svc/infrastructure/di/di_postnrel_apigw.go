package di_postnrel_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	handler_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/handler"
	client_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/client"
	router_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/router"
	"github.com/gofiber/fiber/v2"
)

func InitPostNrelClient(app *fiber.App, config *config_apigw.Config, middleware *middleware_auth_apigw.Middleware) error {

	client, err := client_postnrel_apigw.InitPostNrelClient(config)
	if err != nil {
		return err
	}

	postHandler := handler_postnrel_apigw.NewPostHandler(client)
	relationHandler := handler_postnrel_apigw.NewRelationHandler(client)
	commentHandler := handler_postnrel_apigw.NewCommentHandler(client)

	router_postnrel_apigw.PostNrelUserRoutes(app, postHandler, middleware, relationHandler,commentHandler)

	return nil
}
