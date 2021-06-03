package user

import (
	"context"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user/model"
	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/service"
)

type userService struct {
}

func NewUserService() service.UserService {
	return &userService{}
}

func (u *userService) GetUser(ctx context.Context, id int) model.User {
	return model.User{Id: 1, Name: "Gaurav", Role: "Admin"}
}

func (u *userService) GetAllUser(ctx context.Context) []model.User {
	users := make([]model.User, 0)
	return users
}
