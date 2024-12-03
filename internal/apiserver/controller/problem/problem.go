package problem

import (
	"github.com/ividernvi/algohub-forum/internal/apiserver/service"
	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
)

type ProbController struct {
	ProvSrv service.Service
}

func NewProblemController(store store.Factory) *ProbController {
	return &ProbController{
		ProvSrv: service.NewService(store),
	}
}
