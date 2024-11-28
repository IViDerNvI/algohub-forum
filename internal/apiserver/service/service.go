package service

import "github.com/ividernvi/algohub-forum/internal/apiserver/store"

type Service interface {
	Users() UserSrv
	Posts() PostSrv
	Likes() LikeSrv
	Tokens() TokenSrv
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

func (srv *service) Likes() LikeSrv {
	return newLikeSrv(srv)
}

func (srv *service) Tokens() TokenSrv {
	return newTokenSrc(srv)
}
