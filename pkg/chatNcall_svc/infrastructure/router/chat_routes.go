package router_chatNcallSvc_apigw

import (
	"fmt"

	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	handler_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/handler"
	responsemodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/responsemodels"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func ChatNcallRoutes(app *fiber.App,
	webSocHandler *handler_chatNcallSvc_apigw.ChatWebSocHandler,
	middleware *middleware_auth_apigw.Middleware) {

	app.Use(middleware.UserAuthorizationMiddleware)
	{
		chatManagement := app.Group("/chat")
		{
			chatManagement.Get("/onetoonechats/:recipientid", webSocHandler.GetOneToOneChats)
			chatManagement.Get("/recentchatprofiles/", webSocHandler.GetrecentchatprofileDetails)

			chatManagement.Use(HttptoWsConnectionUpgrader)
			{
				chatManagement.Get("/ws", webSocHandler.WsConnection)

			}

		}

	}
}

func HttptoWsConnectionUpgrader(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}

	fmt.Println("-------------websocket.IsWebSocketUpgrade(ctx)------------,returned false:::::::::::::")
	return ctx.Status(fiber.ErrUpgradeRequired.Code).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.ErrUpgradeRequired.Code,
			Message:    "requires websocket connection",
		})
}
