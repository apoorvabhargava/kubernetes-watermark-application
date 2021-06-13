package user

import (
	"context"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user/model"
	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/repository"
)

type userRepository struct {
	//conn connection object
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) GetUser(ctx context.Context, id int) model.User {
	return model.User{Id: 1, Name: "TestUser", Role: "Admin"}
}

func (u *userRepository) GetAllUser(ctx context.Context) []model.User {
	allUser := make([]model.User, 1)
	allUser = append(allUser, model.User{Id: 1, Name: "TestUser", Role: "Admin"})
	return allUser
}
