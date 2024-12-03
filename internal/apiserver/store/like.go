package store

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
)

type LikeStore interface {
	Create(ctx context.Context, Like model.Like, opt model.CreateOptions) error
	Update(ctx context.Context, Like model.Like, opt model.UpdateOptions) error
	GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Like, error)
	DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error
	List(ctx context.Context, title string, opt model.ListOptions) (*model.LikeList, error)
}
