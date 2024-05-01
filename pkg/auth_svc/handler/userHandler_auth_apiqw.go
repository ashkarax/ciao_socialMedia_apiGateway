package handler_auth_apigw

import (
	"context"
	"fmt"
	"time"

	requestmodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/request_models"
	responsemodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/response_models"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/pb"
	middleware_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Client     pb.AuthServiceClient
	Middleware *middleware_auth_apigw.Middleware
}

func NewAuthUserHandler(client *pb.AuthServiceClient, middleware *middleware_auth_apigw.Middleware) *UserHandler {
	return &UserHandler{Client: *client,
		Middleware: middleware,
	}
}

func (svc *UserHandler) UserSignUp(ctx *fiber.Ctx) error {

	var userSignupData requestmodels_auth_apigw.UserSignUpReq
	var resSignUp responsemodels_auth_apigw.SignupData

	if err := ctx.BodyParser(&userSignupData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "signup failed(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(userSignupData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					resSignUp.Name = "should be a valid Name. "
				case "UserName":
					resSignUp.UserName = "should be a valid username. "
				case "Email":
					resSignUp.Email = "should be a valid email address. "
				case "Password":
					resSignUp.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resSignUp.ConfirmPassword = "should match the first password"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "signup failed",
				Data:       resSignUp,
				Error:      "did't fullfill the signup requirement ",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.UserSignUp(context, &pb.SignUpRequest{
		Name:            userSignupData.Name,
		UserName:        userSignupData.UserName,
		Email:           userSignupData.Email,
		Password:        userSignupData.Password,
		ConfirmPassword: userSignupData.ConfirmPassword,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "signup failed",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "signup failed",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "signup success",
			Data:       resp,
			Error:      nil,
		})
}

func (svc *UserHandler) UserOTPVerication(ctx *fiber.Ctx) error {

	var otpData requestmodels_auth_apigw.OtpVerification
	var otpveriRes responsemodels_auth_apigw.OtpVerifResult

	temptoken := ctx.Get("x-temp-token")

	if err := ctx.BodyParser(&otpData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "OTP verification failed(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(otpData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					otpData.Otp = "otp should be a 4 digit number"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "OTP verification failed",
				Data:       otpveriRes,
				Error:      otpveriRes.Otp,
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.UserOTPVerication(context, &pb.RequestOtpVefification{
		TempToken: temptoken,
		Otp:       otpData.Otp,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "OTP verification failed",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "OTP verification failed",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "OTP verification success",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) UserLogin(ctx *fiber.Ctx) error {

	var loginData requestmodels_auth_apigw.UserLoginReq
	var resLogin responsemodels_auth_apigw.UserLoginRes

	if err := ctx.BodyParser(&loginData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "login failed(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(loginData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					resLogin.Email = "Enter a valid email"
				case "Password":
					resLogin.Password = "Password should have four or more digit"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "OTP verification failed",
				Data:       resLogin,
				Error:      err.Error(),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.UserLogin(context, &pb.RequestUserLogin{
		Email:    loginData.Email,
		Password: loginData.Password,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "login failed",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "login failed",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "login success",
			Data:       resp,
			Error:      nil,
		})
}

func (svc *UserHandler) ForgotPasswordRequest(ctx *fiber.Ctx) error {
	var forgotReqData requestmodels_auth_apigw.ForgotPasswordReq
	var resData responsemodels_auth_apigw.ForgotPasswordRes

	if err := ctx.BodyParser(&forgotReqData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(forgotReqData); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			if len(ve) > 0 && ve[0].Field() == "Email" {
				resData.Email = "Enter a valid email"
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request",
				Data:       resData,
				Error:      err.Error(),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.ForgotPasswordRequest(context, &pb.RequestForgotPass{
		Email: forgotReqData.Email,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed request",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed request",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "success",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) ResetPassword(ctx *fiber.Ctx) error {

	temptoken := ctx.Get("x-temp-token")

	var requestData requestmodels_auth_apigw.ForgotPasswordData
	var resData responsemodels_auth_apigw.ForgotPasswordData

	if err := ctx.BodyParser(&requestData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(requestData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Otp":
					resData.Otp = "otp should be a 4 digit number"
				case "Password":
					resData.Password = "Password should have four or more digit"
				case "ConfirmPassword":
					resData.ConfirmPassword = "should match the first password"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed to reset password",
				Data:       resData,
				Error:      err.Error(),
			})

	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.ResetPassword(context, &pb.RequestResetPass{
		Otp:             requestData.Otp,
		Password:        requestData.Password,
		ConfirmPassword: requestData.ConfirmPassword,
		TempToken:       temptoken,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to reset password",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to reset password",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "password reseted successfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) GetUserProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	UserId := fmt.Sprint(userId)

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetUserProfile(context, &pb.RequestUserId{
		UserId: UserId,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to get user profile",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to get user profile",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched user profile successfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) EditUserProfile(ctx *fiber.Ctx) error {
	var editInput requestmodels_auth_apigw.EditUserProfile
	var respEditUsr responsemodels_auth_apigw.EditUserProfileResp

	userId := ctx.Locals("userId")
	editInput.UserId = fmt.Sprint(userId)

	if err := ctx.BodyParser(&editInput); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "failed request(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(editInput)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					respEditUsr.Name = "should be a valid Name. "
				case "UserName":
					respEditUsr.UserName = "should be a valid username. "
				case "Bio":
					respEditUsr.Bio = "Bio can't exceed 60 characters "
				case "Links":
					respEditUsr.Links = "Links can't exceed 20 characters"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't edit user details",
				Data:       respEditUsr,
				Error:      err.Error(),
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.EditUserProfile(context, &pb.RequestEditUserProfile{
		UserId:   editInput.UserId,
		Name:     editInput.Name,
		UserName: editInput.UserName,
		Bio:      editInput.Bio,
		Links:    editInput.Links,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't edit user details",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't edit user details",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "edited user profile successfully",
			Error:      nil,
		})

}
