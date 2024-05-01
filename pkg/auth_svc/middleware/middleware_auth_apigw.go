package middleware_auth_apigw

import (
	"context"
	"errors"
	"fmt"
	"time"

	responsemodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/response_models"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/pb"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	Client pb.AuthServiceClient
}

func NewAuthMiddleware(client *pb.AuthServiceClient) *Middleware {
	return &Middleware{Client: *client}
}

func (m *Middleware) UserAuthorizationMiddleware(ctx *fiber.Ctx) error {
	accessToken := ctx.Get("x-access-token")

	if accessToken == "" || len(accessToken) < 20 {

		return ctx.Status(fiber.StatusUnauthorized).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "error praising access token from request",
				Error:      errors.New("error praising access token from request"),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := m.Client.VerifyAccessToken(context, &pb.RequestVerifyAccess{
		AccessToken: accessToken,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to verify accesstoken",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to verify accesstoken",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	ctx.Locals("userId", resp.UserId)
	return ctx.Next()
}

func (m *Middleware) AdminAuthMiddleware(ctx *fiber.Ctx) error {

	ctx.Status(fiber.StatusOK)
	return ctx.Next()
}

func (m *Middleware) AccessRegenerator(ctx *fiber.Ctx) error {

	accessToken := ctx.Get("x-access-token")
	refreshToken := ctx.Get("x-refresh-token")

	if accessToken == "" || refreshToken == "" || len(accessToken) < 20 || len(refreshToken) < 20 {
		return ctx.Status(fiber.StatusUnauthorized).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "error praising access and refresh token from request",
				Error:      errors.New("error praising access and refresh token from request"),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := m.Client.AccessRegenerator(context, &pb.RequestAccessGenerator{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to generate accesstoken",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to generate accesstoken",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "new access token generated successfully",
			Data:       resp,
			Error:      nil,
		})
}
