package controller

import (
	"auth_service/pkg/model"
	"context"

	"github.com/gofrs/uuid"
)

//GetUserByEmail(ctx context.Context, email string) (*model.User, error)
//GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
//CreateUser(ctx context.Context, user *model.User) error

type mysqlRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}

type redisRepository interface {
	CreateAuth(userId string, td *TokenDetails) error
	DeleteTokens(authD *AccessDetails) error
	FetchAuth(tokenUuid string) (string, error)
	DeleteRefresh(refreshUuid string) error
}

type Controller struct {
	repo  mysqlRepository
	redis redisRepository
}
