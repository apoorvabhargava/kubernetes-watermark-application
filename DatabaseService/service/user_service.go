package service

import (
	"context"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user/model"
)

type UserService interface {
	GetUser(ctx context.Context, id int) model.User
	GetAllUser(ctx context.Context) []model.User
}
