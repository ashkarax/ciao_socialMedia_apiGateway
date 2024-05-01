package client_apigw_auth

import (
	"fmt"

	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/auth_svc/infrastructure/pb"
	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthClient(config *config_apigw.Config) (*pb.AuthServiceClient, error) {
	cc, err := grpc.Dial(config.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewAuthServiceClient(cc)

	return &Client, nil
}
