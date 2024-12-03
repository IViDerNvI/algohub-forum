package like

import (
	"github.com/ividernvi/algohub-forum/internal/apiserver/service"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
)

type LikeController struct {
	Lsrv service.Service
}

func NewLikeController(store store.Factory) *LikeController {
	return &LikeController{
		Lsrv: service.NewService(store),
	}
}
