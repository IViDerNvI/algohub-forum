package service

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type PostSrv interface {
	Create(ctx context.Context, post model.Post, opt model.CreateOptions) error
	Update(ctx context.Context, post model.Post, opt model.UpdateOptions) error
	GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Post, error)
	DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error
	List(ctx context.Context, username string, opt model.ListOptions) (*model.PostList, error)
}

type postSrv struct {
	store store.Factory
}

func newPostSrv(srv *service) *postSrv {
	return &postSrv{srv.store}
}

func (ps *postSrv) Create(ctx context.Context, post model.Post, opt model.CreateOptions) error {
	return ps.store.Posts().Create(ctx, post, opt)
}

func (ps *postSrv) Update(ctx context.Context, post model.Post, opt model.UpdateOptions) error {
	return ps.store.Posts().Update(ctx, post, opt)
}

func (ps *postSrv) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Post, error) {
	return ps.store.Posts().GetById(ctx, id, opt)
}

func (ps *postSrv) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ps.store.Posts().DeleteById(ctx, id, opt)
}

func (ps *postSrv) List(ctx context.Context, posttitle string, opt model.ListOptions) (*model.PostList, error) {
	return ps.store.Posts().List(ctx, posttitle, opt)
}
