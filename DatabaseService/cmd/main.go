package main

import (
	"context"
	"fmt"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/config"
	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user"
)

func main() {
	fmt.Println("Application Start")
	//read the config file to get the configuration details
	LoadConfig()
	userService := user.NewUserService()
	fmt.Println(userService.GetUser(context.Background(), 1))
}

func LoadConfig() {
	config.LoadConfig()
}
