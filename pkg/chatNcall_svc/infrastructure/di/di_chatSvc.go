package di_chatNcallSvc_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	handler_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/handler"
	client_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/client"
	router_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/router"

	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func InitChatNcallClient(app *fiber.App, config *config_apigw.Config, middleware *middleware_auth_apigw.Middleware) error {

	client, err := client_chatNcallSvc_apigw.InitChatNcallClient(config)
	if err != nil {
		return err
	}

	webSocHandler := handler_chatNcallSvc_apigw.NewChatWebSocHandler(client,config)

	router_chatNcallSvc_apigw.ChatNcallRoutes(app, webSocHandler, middleware)

	return nil
}
