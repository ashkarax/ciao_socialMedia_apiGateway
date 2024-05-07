package client_postnrel_apigw

import (
	"fmt"

	config_apigw "github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/config"
	"github.com/ashkarax/ciao_socialMedia_apiGateway/pkg/postNrelation_svc/infrastructure/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitPostNrelClient(config *config_apigw.Config) (*pb.PostNrelServiceClient, error) {
	cc, err := grpc.Dial(config.PostNrelSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewPostNrelServiceClient(cc)

	return &Client, nil
}
