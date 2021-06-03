package main

import (
	"context"
	"fmt"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user"
)

func main() {
	fmt.Println("Application Restart")
	userService := user.NewUserService()
	fmt.Println(userService.GetUser(context.Background(), 1))
}
