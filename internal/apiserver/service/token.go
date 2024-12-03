package service

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type TokenSrv interface {
	Create(ctx context.Context, token model.Token, opt model.CreateOptions) error
	Get(ctx context.Context, userid string, opt model.GetOptions) (*model.Token, error)
	Delete(ctx context.Context, userid string, opt model.DeleteOptions) error
}

type tokenSrv struct {
	store store.Factory
}

func newTokenSrc(srv *service) *tokenSrv {
	return &tokenSrv{
		store: srv.store,
	}
}

func (ts *tokenSrv) Create(ctx context.Context, token model.Token, opt model.CreateOptions) error {
	return ts.store.Tokens().Create(ctx, token, opt)
}

func (ts *tokenSrv) Get(ctx context.Context, userid string, opt model.GetOptions) (*model.Token, error) {
	return ts.store.Tokens().Get(ctx, userid, opt)
}

func (ts *tokenSrv) Delete(ctx context.Context, userid string, opt model.DeleteOptions) error {
	return ts.store.Tokens().Delete(ctx, userid, opt)
}
