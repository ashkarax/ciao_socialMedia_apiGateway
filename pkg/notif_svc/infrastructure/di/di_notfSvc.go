package di_notifSvc_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	handler_notifSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/notif_svc/handler"
	client_notifSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/notif_svc/infrastructure/client"
	router_notifSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/notif_svc/infrastructure/router"
	"github.com/gofiber/fiber/v2"
)

func InitNotificationClient(app *fiber.App, config *config_apigw.Config, middleware *middleware_auth_apigw.Middleware) error {

	client, err := client_notifSvc_apigw.InitNotificationClient(config)
	if err != nil {
		return err
	}

	notifhandler := handler_notifSvc_apigw.NewNotificationHandler(client, config)

	router_notifSvc_apigw.NotificationRoutes(app, notifhandler, middleware)

	return nil
}
