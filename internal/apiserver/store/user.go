package store

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
)

type UserStore interface {
	Create(ctx context.Context, user model.User, opt model.CreateOptions) error
	Update(ctx context.Context, user model.User, opt model.UpdateOptions) error
	GetById(ctx context.Context, id int, opt model.GetOptions) (*model.User, error)
	GetByAccount(ctx context.Context, account string, opt model.GetOptions) (*model.User, error)
	DeleteById(ctx context.Context, id int, opt model.DeleteOptions) error
	DeleteByUsername(ctx context.Context, username string, opt model.DeleteOptions) error
	List(ctx context.Context, username string, opt model.ListOptions) (*model.UserList, error)
}
