package user

import (
	"github.com/ividernvi/algohub-forum/internal/apiserver/service"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
)

type UserController struct {
	Usrv service.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		Usrv: service.NewService(store),
	}
}
