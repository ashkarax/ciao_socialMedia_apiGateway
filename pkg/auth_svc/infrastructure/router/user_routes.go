package router_auth_apigw

import (
	handler_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/handler"
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthUserRoutes(app *fiber.App, userHandler *handler_auth_apigw.UserHandler, middleware *middleware_auth_apigw.Middleware) {

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
			profileManagement.Patch("/edit", userHandler.EditUserProfile)
			profileManagement.Get("/followers", userHandler.GetFollowersDetails)
			profileManagement.Get("/following", userHandler.GetFollowingsDetails)
		}

		exploremanagement := app.Group("/explore")
		{
			exploremanagement.Get("/profile/:userbid", userHandler.GetAnotherUserProfile)

			// searchmanagement := exploremanagement.Group("/search")
			// {
			// 	searchmanagement.Get("/user/:searchtext", userHandler.SearchUser)

			// }
		}

	}

}
