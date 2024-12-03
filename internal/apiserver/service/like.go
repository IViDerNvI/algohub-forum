package service

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type LikeSrv interface {
	Create(ctx context.Context, Like model.Like, opt model.CreateOptions) error
	Update(ctx context.Context, Like model.Like, opt model.UpdateOptions) error
	GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Like, error)
	DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error
	List(ctx context.Context, username string, opt model.ListOptions) (*model.LikeList, error)
}

type likeSrv struct {
	store store.Factory
}

func newLikeSrv(srv *service) LikeSrv {
	return &likeSrv{srv.store}
}

func (ls *likeSrv) Create(ctx context.Context, Like model.Like, opt model.CreateOptions) error {
	return ls.store.Likes().Create(ctx, Like, opt)
}

func (ls *likeSrv) Update(ctx context.Context, Like model.Like, opt model.UpdateOptions) error {
	return ls.store.Likes().Update(ctx, Like, opt)
}

func (ls *likeSrv) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Like, error) {
	return ls.store.Likes().GetById(ctx, id, opt)
}

func (ls *likeSrv) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ls.store.Likes().DeleteById(ctx, id, opt)
}

func (ls *likeSrv) List(ctx context.Context, Liketitle string, opt model.ListOptions) (*model.LikeList, error) {
	return ls.store.Likes().List(ctx, Liketitle, opt)
}
