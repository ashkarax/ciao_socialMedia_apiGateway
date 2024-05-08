package handler_postnrel_apigw

import (
	"context"
	"fmt"
	"time"

	requestmodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/requestmodels"
	responsemodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/pb"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	Client pb.PostNrelServiceClient
}

func NewCommentHandler(client *pb.PostNrelServiceClient) *CommentHandler {
	return &CommentHandler{Client: *client}
}

func (svc *CommentHandler) AddComment(ctx *fiber.Ctx) error {

	userId := ctx.Locals("userId")

	var commentInput requestmodels_postnrel_apigw.CommentRequest
	var commentOut responsemodels_postnrel_apigw.CommentResponse

	commentInput.UserId = fmt.Sprint(userId)

	if err := ctx.BodyParser(&commentInput); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add comment(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(commentInput)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "UserId":
					commentOut.UserID = "no userId found in request"
				case "CommentText":
					commentOut.CommentText = "should contain less than 30 letters"
				case "PostId":
					commentOut.PostID = "postId should be a number"
				case "ParentCommentId":
					commentOut.ParentCommentID = "ParentCommentId should be a number"
				}
			}
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't add comment",
					Data:       commentOut,
					Error:      err.Error(),
				})
		}
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.AddComment(context, &pb.RequestAddComment{
		UserId:          commentInput.UserId,
		PostId:          uint64(commentInput.PostId),
		CommentText:     commentInput.CommentText,
		ParentCommentId: uint64(commentInput.ParentCommentId),
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't add comment",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't add comment",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "comment added succesfully",
		})
}

func (svc *CommentHandler) DeleteComment(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	commentId := ctx.Params("commentid")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.DeleteComment(context, &pb.RequestCommentDelete{
		UserId:    fmt.Sprint(userId),
		CommentId: commentId,
	})
	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't delete comment",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't delete comment",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "comment deleted succesfully",
		})

}

func (svc *CommentHandler) EditComment(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	var input requestmodels_postnrel_apigw.CommentEditRequest
	var output responsemodels_postnrel_apigw.CommentEditResponse

	input.UserId = fmt.Sprint(userId)

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't edit comment(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(input)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "UserId":
					output.UserId = "no userId found in request"
				case "CommentText":
					output.CommentText = "should contain less than 30 letters"
				case "CommentId":
					output.CommentId = "CommentId should be a number"

				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't edit comment",
				Data:       output,
				Error:      err.Error(),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.EditComment(context, &pb.RequestEditComment{
		CommentId:   uint64(input.CommentId),
		UserId:      input.UserId,
		CommentText: input.CommentText,
	})
	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't edit comment",
				Error:      err.Error(),
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't edit comment",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "comment edited succesfully",
		})
}

func (svc *CommentHandler) FetchPostComments(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	postId := ctx.Params("postid")
	limit, offset := ctx.Query("limit", "5"), ctx.Query("offset", "0")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.FetchPostComments(context, &pb.RequestFetchComments{
		UserId: fmt.Sprint(userId),
		PostId: postId,
		Limit:  limit,
		OffSet: offset,
	})
	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't fetch post comments",
				Error:      err.Error(),
			})
	}
	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't fetch post comments",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "post comments fetched succesfully",
			Data:       resp,
		})

}
