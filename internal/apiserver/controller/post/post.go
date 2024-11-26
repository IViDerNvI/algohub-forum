package post

import (
	"github.com/ividernvi/algohub-forum/internal/apiserver/service"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
)

type PostController struct {
	Psrv service.Service
}

func NewPostController(store store.Factory) *PostController {
	return &PostController{
		Psrv: service.NewService(store),
	}
}
