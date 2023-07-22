package main

import (
	"fmt"
	"log"
	"net"

	"go-grpc-product-svc/pkg/config"
	"go-grpc-product-svc/pkg/db"
	pb "go-grpc-product-svc/pkg/pb"
	services "go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println()
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println(c.DBUrl)

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed at listing:", err)
	}

	fmt.Println("[product-svc] on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
