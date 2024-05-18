package handler_chatNcallSvc_apigw

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	requestmodels_chatNcallSvc_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/models/request_models"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/pb"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type ChatWebSocHandler struct {
	Client *pb.ChatNCallServiceClient
}

func NewChatWebSocHandler(client *pb.ChatNCallServiceClient) *ChatWebSocHandler {
	return &ChatWebSocHandler{Client: client}
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

			switch messageModel.Type {
			case "OneToOne":
				OnetoOneMessage(&messageModel)
			case "OneToMany":
				OnetoMany(&messageModel)
			default:
				sendErrMessageWS(userIdStr, errors.New("message Type should be OneToOne or OneToMany,no other types allowed"))
			}
		}
	})(ctx)

	return nil
}

func OnetoOneMessage(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) {
	fmt.Println(msgModel)

	msgModel.Status = "pending"
	msgModel.Timestamp = time.Now()

	conn, ok := UserSocketMap[msgModel.RecipientID]
	if ok {
		data := MarshalStructJson(msgModel)
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("error sending to recipient", err)
		}
		msgModel.Status = "send"
	}

	fmt.Println("did not find reciever id in SocketMap")
}

func OnetoMany(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) {

}

func sendErrMessageWS(userid string, err error) {
	conn, ok := UserSocketMap[userid]
	if ok {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
}

func MarshalStructJson(msgModel *requestmodels_chatNcallSvc_apigw.MessageRequest) []byte {
	data, err := json.Marshal(msgModel)
	if err != nil {
		sendErrMessageWS(msgModel.SenderID, err)
	}

	return data
}
