package repository

import (
	"context"

	"github.com/apoorvabhargava/kubernetes-watermark-application/database-service/pkg/domain/user/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) model.User
	GetAllUser(ctx context.Context) []model.User
}
