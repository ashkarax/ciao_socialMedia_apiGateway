package client_chatNcallSvc_apigw

import (
	"fmt"

	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/chatNcall_svc/infrastructure/pb"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitChatNcallClient(config *config_apigw.Config) (*pb.ChatNCallServiceClient, error) {
	cc, err := grpc.Dial(config.ChatSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewChatNCallServiceClient(cc)

	return &Client, nil
}
