package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order"
	"go-grpc-api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := gin.Default()

	authSvc := *auth.RegisterRoutes(app, &c)
	product.RegisterRoutes(app, &c, &authSvc)
	order.RegisterRoutes(app, &c, &authSvc)

	app.Run(c.Port)
}
