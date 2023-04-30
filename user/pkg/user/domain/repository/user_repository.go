package repository

import (
	"context"

	"github.com/microservice-ec-site/user/pkg/user/domain/entity"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	ListUsersByIDs(ctx context.Context, ids []string) ([]*entity.User, error)
	Insert(ctx context.Context, input *entity.CreateUserInput) (*entity.User, error)
	Update(ctx context.Context, input *entity.UpdateUserInput) (*entity.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}
