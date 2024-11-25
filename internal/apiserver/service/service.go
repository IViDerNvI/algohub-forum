package service

import "github.com/ividernvi/algohub-forum/internal/apiserver/store"

type Service interface {
	Users() UserSrv
	// PostSrv() PostSrc
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{store: store}
}

func (srv *service) Users() UserSrv {
	return newUserSrv(srv)
}
