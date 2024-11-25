package service

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type UserSrv interface {
	Create(ctx context.Context, user model.User, opt model.CreateOptions) error
	Update(ctx context.Context, user model.User, opt model.UpdateOptions) error
	GetById(ctx context.Context, id int, opt model.GetOptions) (*model.User, error)
	GetByAccount(ctx context.Context, account string, opt model.GetOptions) (*model.User, error)
	DeleteById(ctx context.Context, id int, opt model.DeleteOptions) error
	DeleteByUsername(ctx context.Context, username string, opt model.DeleteOptions) error
	List(ctx context.Context, username string, opt model.ListOptions) (*model.UserList, error)
}

type userSrv struct {
	store store.Factory
}

func newUserSrv(srv *service) *userSrv {
	return &userSrv{srv.store}
}

func (us *userSrv) Create(ctx context.Context, user model.User, opt model.CreateOptions) error {
	return us.store.Users().Create(ctx, user, opt)
}

func (us *userSrv) Update(ctx context.Context, user model.User, opt model.UpdateOptions) error {
	return us.store.Users().Update(ctx, user, opt)
}

func (us *userSrv) GetById(ctx context.Context, id int, opt model.GetOptions) (*model.User, error) {
	return us.store.Users().GetById(ctx, id, opt)
}

func (us *userSrv) GetByAccount(ctx context.Context, account string, opt model.GetOptions) (*model.User, error) {
	return us.store.Users().GetByAccount(ctx, account, opt)
}

func (us *userSrv) DeleteById(ctx context.Context, id int, opt model.DeleteOptions) error {
	return us.store.Users().DeleteById(ctx, id, opt)
}

func (us *userSrv) DeleteByUsername(ctx context.Context, username string, opt model.DeleteOptions) error {
	return us.store.Users().DeleteByUsername(ctx, username, opt)
}

func (us *userSrv) List(ctx context.Context, username string, opt model.ListOptions) (*model.UserList, error) {
	return us.store.Users().List(ctx, username, opt)
}
