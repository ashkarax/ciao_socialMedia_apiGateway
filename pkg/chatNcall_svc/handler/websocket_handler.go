package handler_chatNcallSvc_apigw

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	requestmodels_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/models/request_models"
	responsemodels_chatNcall "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/models/response_models"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/pb"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	responsemodels_postnrel_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/models/responsemodels"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ChatWebSocHandler struct {
	Client      pb.ChatNCallServiceClient
	LocationInd *time.Location
	Config      *config_apigw.Config
}

func NewChatWebSocHandler(client *pb.ChatNCallServiceClient, config *config_apigw.Config) *ChatWebSocHandler {
	locationInd, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("----------------------error fetching location:", err)
	}
	return &ChatWebSocHandler{Client: *client,
		LocationInd: locationInd,
		Config:      config,
	}
}

var UserSocketMap = make(map[string]*websocket.Conn)

func (svc *ChatWebSocHandler) WsConnection(ctx *fiber.Ctx) error {
	websocket.New(func(conn *websocket.Conn) {
		var messageModel requestmodels_chatNcallSvc_apigw.MessageRequest
		var (
			msg []byte
			err error
		)

		userId := conn.Locals("userId")
		userIdStr := fmt.Sprint(userId)

		UserSocketMap[userIdStr] = conn

		defer conn.Close()
		defer delete(UserSocketMap, userIdStr)

		for {
			if _, msg, err = conn.ReadMessage(); err != nil {
				log.Println("read:", err)
				sendErrMessageWS(userIdStr, err)
				break
			}
			err = json.Unmarshal(msg, &messageModel)
			if err != nil {
				log.Println("read:", err)
				sendErrMessageWS(userIdStr, err)
				break
			}
			messageModel.SenderID = userIdStr
			messageModel.TimeStamp = time.Now()

			validate := validator.New(validator.WithRequiredStructEnabled())
			err = validate.Struct(messageModel)
			if err != nil {
				if ve, ok := err.(validator.ValidationErrors); ok {
					for _, e := range ve {
						switch e.Field() {
						case "Type":
							sendErrMessageWS(userIdStr, errors.New("no Type found in input"))
						}
					}
				}
				break
			}
			switch messageModel.Type {
			case "OneToOne":
				svc.OnetoOneMessage(&messageModel)
			case "TypingStatus":
				svc.TypingStatus(&messageModel)
			case "OneToMany":
				svc.OnetoMany(&messageModel)
			default:
				sendErrMessageWS(userIdStr, errors.New("message Type should be OneToOne,OneToMany,DeleteMessage,UpdateSeenStatus or TypingStatus ,no other types allowed"))
			}
		}
	})(ctx)

	return nil
}

func (svc *ChatWebSocHandler) TypingStatus(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) {
	if msgModel.RecipientID == "" {
		sendErrMessageWS(msgModel.SenderID, errors.New("no RecipientID found in input"))
	}
	var MsgModel requestmodels_chatNcallSvc_apigw.TypingStatusRequest

	MsgModel.SenderID = msgModel.SenderID
	MsgModel.RecipientID = msgModel.RecipientID
	MsgModel.Type = msgModel.Type
	MsgModel.TypingStat = msgModel.TypingStat

	conn, ok := UserSocketMap[MsgModel.RecipientID]
	if ok {
		data, err := MarshalStructJson(MsgModel)
		if err != nil {
			sendErrMessageWS(MsgModel.SenderID, err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, *data)
		if err != nil {
			fmt.Println("error sending to recipient", err)
			return
		}
	}
}

func (svc *ChatWebSocHandler) OnetoOneMessage(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) {
	if msgModel.RecipientID == "" {
		sendErrMessageWS(msgModel.SenderID, errors.New("no RecipientID found in input"))
	}
	if msgModel.Content == "" || len(msgModel.Content) > 100 {
		sendErrMessageWS(msgModel.SenderID, errors.New("message content should be less than 100 words "))
		return
	}

	var OneToOneMsgModel requestmodels_chatNcallSvc_apigw.OnetoOneMessageRequest

	OneToOneMsgModel.SenderID = msgModel.SenderID
	OneToOneMsgModel.RecipientID = msgModel.RecipientID
	OneToOneMsgModel.Content = msgModel.Content
	OneToOneMsgModel.TimeStamp = msgModel.TimeStamp
	OneToOneMsgModel.Status = "pending"
	OneToOneMsgModel.Type = msgModel.Type

	conn, ok := UserSocketMap[OneToOneMsgModel.RecipientID]
	if ok {
		OneToOneMsgModel.TimeStamp = OneToOneMsgModel.TimeStamp.In(svc.LocationInd)
		data, err := MarshalStructJson(OneToOneMsgModel)
		if err != nil {
			sendErrMessageWS(OneToOneMsgModel.SenderID, err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, *data)
		if err != nil {
			fmt.Println("error sending to recipient", err)
			return
		}
		OneToOneMsgModel.Status = "send"
	}

	fmt.Println("------check status is pending or send--------", OneToOneMsgModel)
	fmt.Println("Adding to kafkaproducer for transporting to service and storing")
	svc.KafkaProducerUpdateOneToOneMessage(&OneToOneMsgModel)
}

func (svc *ChatWebSocHandler) OnetoMany(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) {
	if msgModel.GroupID == "" {
		sendErrMessageWS(msgModel.SenderID, errors.New("no GroupID found in input"))
		return
	}
	if msgModel.Content == "" || len(msgModel.Content) > 100 {
		sendErrMessageWS(msgModel.SenderID, errors.New("message content should be less than 100 words "))
		return
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := svc.Client.GetGroupMembersInfo(context, &pb.RequestGroupMembersInfo{GroupID: msgModel.GroupID})
	if err != nil {
		log.Println("-----from handler:OnetoMany() chatNcall service down while calling GetGroupMembersInfo()---")
		sendErrMessageWS(msgModel.SenderID, err)
		return
	}
	if resp.ErrorMessage != "" {
		sendErrMessageWS(msgModel.SenderID, errors.New(resp.ErrorMessage))
		return
	}

	var OneToManyMsgModel requestmodels_chatNcallSvc_apigw.OnetoManyMessageRequest
	OneToManyMsgModel.SenderID = msgModel.SenderID
	OneToManyMsgModel.GroupID = msgModel.GroupID
	OneToManyMsgModel.Content = msgModel.Content
	OneToManyMsgModel.TimeStamp = msgModel.TimeStamp
	OneToManyMsgModel.Status = "pending"
	OneToManyMsgModel.Type = msgModel.Type

	for i := range resp.GroupMembers {
		if (resp.GroupMembers)[i] == msgModel.SenderID {
			continue
		}
		conn, ok := UserSocketMap[(resp.GroupMembers[i])]
		if ok {
			OneToManyMsgModel.TimeStamp = OneToManyMsgModel.TimeStamp.In(svc.LocationInd)
			data, err := MarshalStructJson(OneToManyMsgModel)
			if err != nil {
				sendErrMessageWS(OneToManyMsgModel.SenderID, err)
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, *data)
			if err != nil {
				fmt.Println("error sending to recipient", err)
				return
			}
			OneToManyMsgModel.Status = "send"
		}
	}

	fmt.Println("------check status is pending or send--------", OneToManyMsgModel)
	fmt.Println("Adding to kafkaproducer for transporting to service and storing")
	svc.KafkaProducerUpdateOneToManyMessage(&OneToManyMsgModel)

}

func (svc *ChatWebSocHandler) KafkaProducerUpdateOneToOneMessage(message *requestmodels_chatNcallSvc_apigw.OnetoOneMessageRequest) error {
	fmt.Println("---------------to KafkaProducerUpdateOneToOneMessage:", *message)

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{svc.Config.KafkaPort}, configs)
	if err != nil {
		return err
	}

	msgJson, _ := MarshalStructJson(message)

	msg := &sarama.ProducerMessage{Topic: svc.Config.KafkaTopicOneToOne,
		Key:   sarama.StringEncoder(message.RecipientID),
		Value: sarama.StringEncoder(*msgJson)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err sending message to kafkaproducer ", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func (svc *ChatWebSocHandler) KafkaProducerUpdateOneToManyMessage(message *requestmodels_chatNcallSvc_apigw.OnetoManyMessageRequest) error {
	fmt.Println("---------------to KafkaProducerUpdateOneToManyMessage:", *message)

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{svc.Config.KafkaPort}, configs)
	if err != nil {
		return err
	}

	msgJson, _ := MarshalStructJson(message)

	msg := &sarama.ProducerMessage{Topic: svc.Config.KafkaTopicOneToMany,
		Key:   sarama.StringEncoder(message.GroupID),
		Value: sarama.StringEncoder(*msgJson)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err sending message to kafkaproducer ", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func (svc *ChatWebSocHandler) GetOneToOneChats(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	limit, offset := ctx.Query("limit", "12"), ctx.Query("offset", "0")

	recepientId := ctx.Params("recipientid")
	if recepientId == "" {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't get chat(possible-reason:no input)",
				Error:      "no recepientid found in request",
			})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := svc.Client.GetOneToOneChats(context, &pb.RequestUserOneToOneChat{
		SenderID:   fmt.Sprint(userId),
		RecieverID: recepientId,
		Limit:      limit,
		Offset:     offset,
	})

	if err != nil {
		fmt.Println("----------chatNcall service down--------")
		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't get chat",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't get chat",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Data:       resp,
			Message:    "chat fetched succesfully",
		})

}

func (svc *ChatWebSocHandler) GetrecentchatprofileDetails(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	limit, offset := ctx.Query("limit", "12"), ctx.Query("offset", "0")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := svc.Client.GetRecentChatProfiles(context, &pb.RequestRecentChatProfiles{
		SenderID: fmt.Sprint(userId),
		Limit:    limit,
		Offset:   offset,
	})

	if err != nil {
		fmt.Println("----------chatNcall service down--------,err:", err)
		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't get recent chat profiles",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		fmt.Println("-----------------------", resp.ErrorMessage)
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't get recent chat profiles",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Data:       resp,
			Message:    "recent chat profiles fetched succesfully",
		})

}

func (svc *ChatWebSocHandler) CreateNewGroup(ctx *fiber.Ctx) error {
	var newGroupData requestmodels_chatNcallSvc_apigw.NewGroupInfo
	userId := ctx.Locals("userId")
	userIdInt, _ := strconv.Atoi(fmt.Sprint(userId))

	if err := ctx.BodyParser(&newGroupData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't create Group(possible-reason:invalid/no json input)",
				Error:      err.Error(),
			})
	}

	newGroupData.CreatorID = fmt.Sprint(userId)
	found := false
	for _, member := range newGroupData.GroupMembers {
		if member == uint64(userIdInt) {
			found = true
			break
		}
	}
	if !found {
		newGroupData.GroupMembers = append(newGroupData.GroupMembers, uint64(userIdInt))
	}

	var validationReponse responsemodels_chatNcall.NewGroupInfo

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(newGroupData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "GroupName":
					validationReponse.GroupName = "should contain less than 20 letters"
				case "GroupMembers":
					validationReponse.GroupMembers = "Should be unique,maximum 12 members and id should be a number"
				}
			}
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't create group",
					Data:       validationReponse,
					Error:      err.Error(),
				})
		}
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.CreateNewGroup(context, &pb.RequestNewGroup{
		GroupName:    newGroupData.GroupName,
		GroupMembers: newGroupData.GroupMembers,
		CreatorID:    newGroupData.CreatorID,
		CreatedAt:    fmt.Sprint(time.Now()),
	})

	if err != nil {
		log.Println("-----error: from handler:CreateNewGroup(),chatNcall service down while calling CreateNewGroup()")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't create group",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't create group",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "group created succesfully",
		})

}

func (svc *ChatWebSocHandler) GetUserGroupsAndLastMessage(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")
	limit, offset := ctx.Query("limit", "12"), ctx.Query("offset", "0")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetUserGroupChatSummary(context, &pb.RequestGroupChatSummary{
		UserID: fmt.Sprint(userId),
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		log.Println("-----error: from handler:GetUserGroupsAndLastMessage(),chatNcall service down while calling GetUserGroupChatSummary()")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to fetch groupchat summary",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to fetch groupchat summary",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Data:       resp.SingleEntity,
			Message:    "groupchat summary fetched succesfully",
		})

}

func sendErrMessageWS(userid string, err error) {
	conn, ok := UserSocketMap[userid]
	if ok {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
}

func MarshalStructJson(msgModel interface{}) (*[]byte, error) {
	data, err := json.Marshal(msgModel)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (svc *ChatWebSocHandler) GetGroupChats(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	groupId := ctx.Params("groupid")
	limit, offset := ctx.Query("limit", "12"), ctx.Query("offset", "0")

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.GetOneToManyChats(context, &pb.RequestGetOneToManyChats{
		UserID:  fmt.Sprint(userId),
		GroupID: groupId,
		Limit:   limit,
		Offset:  offset,
	})

	if err != nil {
		log.Println("-----error: from handler:GetGroupChats(),chatNcall service down while calling GetOneToManyChats()")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "failed to fetch groupchat",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "failed to fetch groupchat",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Data:       resp.Chat,
			Message:    "groupchat fetched succesfully",
		})
}

func (svc *ChatWebSocHandler) AddMembersToGroup(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	var addMembersInput requestmodels_chatNcallSvc_apigw.AddNewMembersToGroup

	if err := ctx.BodyParser(&addMembersInput); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't add new GroupMembers(possible-reason:invalid/no json input)",
				Error:      err.Error(),
			})
	}

	var validationReponse responsemodels_chatNcall.AddNewMembersToGroupResponse
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(addMembersInput)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "GroupID":
					validationReponse.GroupID = "should contain less than 35 characters"
				case "GroupMembers":
					validationReponse.GroupMembers = "Should be unique,maximum 12 members and id should be a number"
				}
			}
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't add new GroupMembers",
					Data:       validationReponse,
					Error:      err.Error(),
				})
		}
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.AddMembersToGroup(context, &pb.RequestAddGroupMembers{
		UserID:    fmt.Sprint(userId),
		GroupID:   addMembersInput.GroupID,
		MemberIDs: addMembersInput.GroupMembers,
	})

	if err != nil {
		log.Println("-----error: from handler:AddMembersToGroup(),chatNcall service down while calling AddMembersToGroup()")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't add new GroupMembers",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't add new GroupMembers",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "new groupmembers added succesfully",
		})

}

func (svc *ChatWebSocHandler) RemoveAMemberFromGroup(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId")

	var inputData requestmodels_chatNcallSvc_apigw.RemoveMemberFromGroup

	if err := ctx.BodyParser(&inputData); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.ErrBadRequest.Code,
				Message:    "can't remove GroupMembers(possible-reason:invalid/no json input)",
				Error:      err.Error(),
			})
	}

	var validationReponse responsemodels_chatNcall.RemoveMemberFromGroup
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(validationReponse)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "GroupID":
					validationReponse.GroupID = "should contain less than 35 characters"
				case "MemberID":
					validationReponse.MemberID = "Should be a valid id"
				}
			}
			return ctx.Status(fiber.ErrBadRequest.Code).
				JSON(responsemodels_postnrel_apigw.CommonResponse{
					StatusCode: fiber.ErrBadRequest.Code,
					Message:    "can't remove GroupMember",
					Data:       validationReponse,
					Error:      err.Error(),
				})
		}
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := svc.Client.RemoveMemberFromGroup(context, &pb.RequestRemoveGroupMember{
		UserID:   fmt.Sprint(userId),
		GroupID:  inputData.GroupID,
		MemberID: inputData.MemberID,
	})

	if err != nil {
		log.Println("-----error: from handler:RemoveAMemberFromGroup(),chatNcall service down while calling RemoveMemberFromGroup()")

		return ctx.Status(fiber.StatusServiceUnavailable).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusServiceUnavailable,
				Message:    "can't remove GroupMember",
				Error:      err.Error(),
			})
	}

	if resp.ErrorMessage != "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(responsemodels_postnrel_apigw.CommonResponse{
				StatusCode: fiber.StatusBadRequest,
				Message:    "can't remove GroupMember",
				Data:       resp,
				Error:      resp.ErrorMessage,
			})
	}

	return ctx.Status(fiber.StatusOK).
		JSON(responsemodels_postnrel_apigw.CommonResponse{
			StatusCode: fiber.StatusOK,
			Message:    "groupmember removed succesfully",
		})

}
