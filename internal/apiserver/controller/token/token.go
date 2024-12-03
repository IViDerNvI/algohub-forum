package token

import (
	"github.com/ividernvi/algohub-forum/internal/apiserver/service"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
)

type TokenController struct {
	Tsrv service.Service
}

func NewTokenController(store store.Factory) *TokenController {
	return &TokenController{
		Tsrv: service.NewService(store),
	}
}
