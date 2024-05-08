package router_postnrel_apigw

import (
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	handler_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/handler"
	"github.com/gofiber/fiber/v2"
)

func PostNrelUserRoutes(app *fiber.App,
	postHandler *handler_postnrel_apigw.PostHandler,
	middleware *middleware_auth_apigw.Middleware,
	relationHandler *handler_postnrel_apigw.RelationHandler,
	commentHandler *handler_postnrel_apigw.CommentHandler) {

	app.Use(middleware.UserAuthorizationMiddleware)
	{
		postManagement := app.Group("/post")
		{
			postManagement.Post("/", postHandler.AddNewPost)
			postManagement.Get("/", postHandler.GetAllPostByUser)
			postManagement.Delete("/:postid", postHandler.DeletePost)
			postManagement.Patch("/", postHandler.EditPost)

			postManagement.Get("/userrelatedposts", postHandler.GetAllRelatedPostsForHomeScreen)


			likemanagement := postManagement.Group("/like")
			{
				likemanagement.Post("/:postid", postHandler.LikePost)
				likemanagement.Delete("/:postid", postHandler.UnLikePost)
			}
			commentManagement := postManagement.Group("/comment")
			{
				commentManagement.Get("/:postid", commentHandler.FetchPostComments)
				commentManagement.Post("/", commentHandler.AddComment)
				commentManagement.Delete("/:commentid", commentHandler.DeleteComment)
				commentManagement.Patch("/", commentHandler.EditComment)
			}

		}
		followRelationshipManagement := app.Group("/relation")
		{
			followRelationshipManagement.Post("/follow/:followingid", relationHandler.Follow)
			followRelationshipManagement.Delete("/unfollow/:unfollowingid", relationHandler.UnFollow)

		}
		exploremanagement := app.Group("/explore")
		{
			exploremanagement.Get("/", postHandler.GetMostLovedPostsFromGlobalUser)

		}

	}
}
