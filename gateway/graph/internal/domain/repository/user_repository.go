package repository

import (
	"context"

	"github.com/microservice-ec-site/graph/internal/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	ListUsersByIDs(ctx context.Context, ids []int) ([]*model.User, error)
	Insert(ctx context.Context, input *model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input *model.UpdateUserInput) (*model.User, error)
	Delete(ctx context.Context, id int) (bool, error)
}
