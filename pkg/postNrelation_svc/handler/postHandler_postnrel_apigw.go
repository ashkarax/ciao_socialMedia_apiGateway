package handler_postnrel_apigw

import (
	"context"
	"fmt"
	"log"
	"time"

	requestmodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/requestmodels"
	responsemodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/responsemodels"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/pb"
	byteconverter_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/utils/byte_converter"
	mediafileformatchecker_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/utils/mediaFileFormatChecker"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	Client pb.PostNrelServiceClient
}

func NewPostHandler(client *pb.PostNrelServiceClient) *PostHandler {
	return &PostHandler{Client: *client}
}

func (svc *PostHandler) AddNewPost(ctx *fiber.Ctx) error {

	var postData requestmodels_postnrel_apigw.AddPostData
	var respPostData responsemodels_postnrel_apigw.AddPostResp

	userId := ctx.Locals("userId")
	postData.UserId = fmt.Sprint(userId)

	if err := ctx.BodyParser(&postData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add post(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	//fiber's ctx.BodyParser can't parse files(*multipart.FileHeader),
	//so we have to manually access the Multipart form and read the files form it.
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["media"]
	postData.Media = append(postData.Media, files...)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(postData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Caption":
					respPostData.Caption = "should contain less than 60 letters"
				case "UserId":
					respPostData.UserId = "No userId got"
				case "Media":
					respPostData.Media = "you can't add a post without a image/video"
				}
			}
		}
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add post",
				Data:       respPostData,
				Error:      err.Error(),
			})
	}

	numFiles := len(postData.Media)
	if numFiles < 1 || numFiles > 5 {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add post",
				Data:       nil,
				Error:      "you can only add 5 image/video in a post",
			})
	}

	for _, media := range postData.Media {
		if media.Size > 5*1024*1024 { // 5 MB limit
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't add post",
					Data:       nil,
					Error:      "yfile size exceeds the limit (5MB)",
				})
		}
	}

	var mediaData []*pb.SingleMedia
	for _, fileHeader := range postData.Media {
		file, err := fileHeader.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		contentType, err := mediafileformatchecker_apigw.MediaFileFormatChecker(file)
		if err != nil {
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't add post",
					Data:       nil,
					Error:      err.Error(),
				})
		}

		content, err := byteconverter_apigw.MultipartFileheaderToBytes(&file)
		if err != nil {
			fmt.Println("-------------byteconverter-down---------")
		}

		mediaData = append(mediaData, &pb.SingleMedia{Media: content, ContentType: *contentType})

	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.AddNewPost(context, &pb.RequestAddPost{
		UserId:  postData.UserId,
		Caption: postData.Caption,
		Media:   mediaData,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't add post",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't add post",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Post added succesfully",
			Data:       resp,
			Error:      nil,
		})
}

func (svc *PostHandler) GetAllPostByUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	limit, offset := ctx.Query("limit", "12"), ctx.Query("offset", "0")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetAllPostByUser(context, &pb.RequestGetAllPosts{
		UserId: fmt.Sprint(userId),
		Limit:  limit,
		OffSet: offset,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't fetch Posts",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't fetch Posts",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Posts fetched succesfully",
			Data:       resp,
			Error:      nil,
		})

}

func (svc *PostHandler) DeletePost(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	postId := ctx.Params("postid")

	if fmt.Sprint(userId) == "" || postId == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't delete post",
				Data:       nil,
				Error:      "no postid found in request",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.DeletePost(context, &pb.RequestDeletePost{
		UserId: fmt.Sprint(userId),
		PostId: postId,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't delete Posts",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't delete Posts",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Post deleted succesfully",
		})

}

func (svc *PostHandler) EditPost(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	var editInput requestmodels_postnrel_apigw.EditPost
	var respPostEdit responsemodels_postnrel_apigw.EditPostResp

	editInput.UserId = fmt.Sprint(userId)

	if err := ctx.BodyParser(&editInput); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add post(possible-reason:no json input)",
				Error:      err.Error(),
			})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(editInput)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Caption":
					respPostEdit.Caption = "should contain less than 60 letters"
				case "UserId":
					respPostEdit.UserId = "No userId got from header"
				case "PostId":
					respPostEdit.PostId = "no postid found in request"

				}
			}
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't edit post",
					Data:       respPostEdit,
					Error:      err.Error(),
				})
		}
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.EditPost(context, &pb.RequestEditPost{
		UserId:  editInput.UserId,
		PostId:  editInput.PostId,
		Caption: editInput.Caption,
	})

	if err != nil {
		fmt.Println("----------postNrel service down--------")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't edit Posts",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't edit Posts",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "Posts edited succesfully",
		})

}
