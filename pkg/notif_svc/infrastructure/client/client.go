package client_notifSvc_apigw

import (
	"fmt"

	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/notif_svc/infrastructure/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitNotificationClient(config *config_apigw.Config) (*pb.NotificationServiceClient, error) {
	cc, err := grpc.Dial(config.NotifSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewNotificationServiceClient(cc)

	return &Client, nil
}
