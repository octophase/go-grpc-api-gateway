package order

import (
	"fmt"

	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
