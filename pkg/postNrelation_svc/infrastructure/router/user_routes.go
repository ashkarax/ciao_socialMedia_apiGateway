package router_postnrel_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	handler_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/handler"
	"github.com/gofiber/fiber/v2"
)

func PostNrelUserRoutes(app *fiber.App,
	postHandler *handler_postnrel_apigw.PostHandler,
	middleware *middleware_auth_apigw.Middleware,
	relationHandler *handler_postnrel_apigw.RelationHandler) {

	app.Use(middleware.UserAuthorizationMiddleware)
	{
		postManagement := app.Group("/post")
		{
			postManagement.Post("/", postHandler.AddNewPost)
			postManagement.Get("/", postHandler.GetAllPostByUser)
			postManagement.Delete("/:postid", postHandler.DeletePost)
			postManagement.Patch("/", postHandler.EditPost)

		}
		followRelationshipManagement := app.Group("/relation")
		{
			followRelationshipManagement.Post("/follow/:followingid", relationHandler.Follow)
			followRelationshipManagement.Delete("/unfollow/:unfollowingid", relationHandler.UnFollow)

		}

	}
}
