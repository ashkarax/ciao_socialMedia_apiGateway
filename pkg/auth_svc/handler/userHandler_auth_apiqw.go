package handler_auth_apigw

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	requestmodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/request_models"
	responsemodels_auth_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/models/response_models"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/pb"
	byteconverter_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/utils/byte_converter"
	mediafileformatchecker_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/utils/mediaFileFormatChecker"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Client pb.AuthServiceClient
	// Middleware *middleware_auth_apigw.Middleware
}

func NewAuthUserHandler(client *pb.AuthServiceClient) *UserHandler {
	return &UserHandler{Client: *client}
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

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetUserProfile(context, &pb.RequestGetUserProfile{
		UserId: fmt.Sprint(userId),
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

	var respStruct responsemodels_auth_apigw.UserProfileA //used to show the zero count of posts,following,followers etc

	intValueuserId, _ := strconv.Atoi(fmt.Sprint(userId))
	uintValueuserId := uint(intValueuserId)

	respStruct.UserId = uintValueuserId
	respStruct.Name = resp.Name
	respStruct.UserName = resp.UserName
	respStruct.Bio = resp.Bio
	respStruct.Links = resp.Links
	respStruct.UserProfileImgURL = resp.ProfileImageURL
	respStruct.PostsCount = uint(resp.PostsCount)
	respStruct.FollowersCount = uint(resp.FollowerCount)
	respStruct.FollowingCount = uint(resp.FollowingCount)

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched user profile successfully",
			Data:       respStruct,
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

func (svc *UserHandler) GetFollowersDetails(ctx *fiber.Ctx) error {

	userId := ctx.Locals("userId")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetFollowersDetails(context, &pb.RequestUserId{UserId: fmt.Sprint(userId)})
	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't fetch followers details",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't fetch followers details",
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched followers details successfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) GetFollowingsDetails(ctx *fiber.Ctx) error {

	userId := ctx.Locals("userId")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetFollowingsDetails(context, &pb.RequestUserId{UserId: fmt.Sprint(userId)})
	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't fetch followings details",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't fetch followings details",
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched followings details successfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *UserHandler) GetAnotherUserProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	UserId := fmt.Sprint(userId)

	userBId := ctx.Params("userbid")

	if fmt.Sprint(userId) == "" || userBId == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't fetch user profile",
				Data:       nil,
				Error:      "no userbid found in request",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetUserProfile(context, &pb.RequestGetUserProfile{
		UserId:  UserId,
		UserBId: userBId,
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

	var respStruct responsemodels_auth_apigw.UserProfileB //used to show the zero count of posts,following,followers etc

	intValueuserBId, _ := strconv.Atoi(userBId)
	uintValueuserBId := uint(intValueuserBId)

	respStruct.UserId = uintValueuserBId
	respStruct.Name = resp.Name
	respStruct.UserName = resp.UserName
	respStruct.Bio = resp.Bio
	respStruct.Links = resp.Links
	respStruct.UserProfileImgURL = resp.ProfileImageURL
	respStruct.PostsCount = uint(resp.PostsCount)
	respStruct.FollowersCount = uint(resp.FollowerCount)
	respStruct.FollowingCount = uint(resp.FollowingCount)
	respStruct.FollowingStatus = resp.FollowingStat

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched user profile successfully",
			Data:       respStruct,
			Error:      nil,
		})

}

func (svc *UserHandler) SearchUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	searchText := ctx.Params("searchtext")
	limit, offset := ctx.Query("limit", "5"), ctx.Query("offset", "0")

	if searchText == "" {
		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to get search result",
				Error:      "enter a valid name or username",
			})
	}

	validSearch := regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`).MatchString
	if len(searchText) > 12 || !validSearch(searchText) {
		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to get search result",
				Error:      "searchtext should contain only less than 12 letters and search input can only contain letters, numbers, spaces, or underscores",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.SearchUser(context, &pb.RequestUserSearch{
		UserId:     fmt.Sprint(userId),
		SearchText: searchText,
		Limit:      limit,
		Offset:     offset,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to get search result",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to get search result",
				Error:      resp.ErrorMessage,
			})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "fetched search result successfully",
			Data:       resp,
			Error:      nil,
		})
}

func (svc *UserHandler) SetProfileImage(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	//fiber's ctx.BodyParser can't parse files(*multipart.FileHeader),
	//so we have to manually access the Multipart form and read the files form it.
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	img := form.File["ProfileImg"]
	if len(img) == 0 || len(img) > 1 {

		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add profile image",
				Error:      "no ProfileImg found in request,you should exactly upload only one img",
			})

	}
	ProfileImg := img[0]

	if ProfileImg.Size > 2*1024*1024 { // 2 MB limit
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add profile image",
				Error:      "ProfileImg size exceeds the limit (2MB)",
			})
	}

	file, err := ProfileImg.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contentType, err := mediafileformatchecker_apigw.ProfileImageFileFormatChecker(file)
	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add profile image",
				Error:      err.Error(),
			})
	}

	content, err := byteconverter_apigw.MultipartFileheaderToBytes(&file)
	if err != nil {
		fmt.Println("-------------byteconverter-down---------")
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.SetUserProfileImage(context, &pb.RequestSetProfileImg{
		UserId:      fmt.Sprint(userId),
		ContentType: *contentType,
		Img:         content,
	})

	if err != nil {
		fmt.Println("----------auth service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't add profile image",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_auth_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't add profile image",
				Error:      resp.ErrorMessage,
			})
	}
	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_auth_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "profile image set successfully",
			Error:      nil,
		})

}
