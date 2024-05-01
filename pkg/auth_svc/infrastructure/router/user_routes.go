package router_auth_apigw

import (
	"net/http"

	handler_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/handler"
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, config *config_apigw.Config, userHandler *handler_auth_apigw.UserHandler, middleware *middleware_auth_apigw.Middleware) {

	app.Use(func(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-Key")
		if apiKey != config.ApiKey {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API key"})
		}
		return c.Next()
	})

	app.Post("/signup", userHandler.UserSignUp)
	app.Post("/verify", userHandler.UserOTPVerication)
	app.Post("/login", userHandler.UserLogin)
	app.Post("/forgotpassword", userHandler.ForgotPasswordRequest)
	app.Patch("/resetpassword", userHandler.ResetPassword)
	app.Get("/accessgenerator", middleware.AccessRegenerator)

	app.Use(middleware.UserAuthorizationMiddleware)
	{
		profileManagement := app.Group("/profile")
		{
			profileManagement.Get("/", userHandler.GetUserProfile)
			profileManagement.Patch("/edit",userHandler.EditUserProfile)
		}
	}

}
