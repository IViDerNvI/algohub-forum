package store

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
)

type PostStore interface {
	Create(ctx context.Context, post model.Post, opt model.CreateOptions) error
	Update(ctx context.Context, post model.Post, opt model.UpdateOptions) error
	GetById(ctx context.Context, id int, opt model.GetOptions) (*model.Post, error)
	DeleteById(ctx context.Context, id int, opt model.DeleteOptions) error
	List(ctx context.Context, opt model.ListOptions) (*model.PostList, error)
}
