package handler_postnrel_apigw

import (
	"context"
	"fmt"
	"time"

	responsemodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/pb"
	"github.com/gofiber/fiber/v2"
)

type RelationHandler struct {
	Client pb.PostNrelServiceClient
}

func NewRelationHandler(client *pb.PostNrelServiceClient) *RelationHandler {
	return &RelationHandler{Client: *client}
}

func (svc *RelationHandler) Follow(ctx *fiber.Ctx) error {
	userid := ctx.Locals("userId")
	userId := fmt.Sprint(userid)

	userBId := ctx.Params("followingid")

	if userBId == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request(possible-reason:no input)",
				Error:      "no userBId (\":followingid\") param found in request.",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.Follow(context, &pb.RequestFollowUnFollow{
		UserId:  userId,
		UserBId: userBId,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to follow userB",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to follow userB",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "followed succesfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *RelationHandler) UnFollow(ctx *fiber.Ctx) error {

	userid := ctx.Locals("userId")
	userId := fmt.Sprint(userid)

	userBId := ctx.Params("unfollowingid")

	if userBId == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request(possible-reason:no input)",
				Error:      "no userBId (\":unfollowingid\") param found in request.",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.UnFollow(context, &pb.RequestFollowUnFollow{
		UserId:  userId,
		UserBId: userBId,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to follow userB",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to follow userB",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "followed succesfully",
			Data:       resp,
			Error:      nil,
		})
}
