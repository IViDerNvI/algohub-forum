package service

import "github.com/ividernvi/algohub-forum/internal/apiserver/store"

type Service interface {
	Users() UserSrv
	Posts() PostSrv
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

func (srv *service) Posts() PostSrv {
	return newPostSrv(srv)
}
