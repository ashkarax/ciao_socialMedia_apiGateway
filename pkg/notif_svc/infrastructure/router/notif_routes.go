package router_notifSvc_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	handler_notifSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/notif_svc/handler"
	"github.com/gofiber/fiber/v2"
)

func NotificationRoutes(app *fiber.App,
	NotifHandler *handler_notifSvc_apigw.NotifHandler,
	middleware *middleware_auth_apigw.Middleware) {

	app.Use(middleware.UserAuthorizationMiddleware)
	{
		notifManger := app.Group("/notification")
		{
			notifManger.Get("/", NotifHandler.GetNotificationsForUser)
		}

	}

}
