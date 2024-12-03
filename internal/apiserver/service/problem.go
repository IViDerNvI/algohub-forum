package service

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/apiserver/store"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type ProbSrv interface {
	Create(ctx context.Context, prob *model.Problem, opt model.CreateOptions) error
	Update(ctx context.Context, prob *model.Problem, opt model.UpdateOptions) error
	GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Problem, error)
	DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error
	List(ctx context.Context, opt model.ListOptions) (*model.ProblemList, error)
}

type probSrv struct {
	store store.Factory
}

func newProbSrv(srv *service) *probSrv {
	return &probSrv{store: srv.store}
}

func (ps *probSrv) Create(ctx context.Context, prob *model.Problem, opt model.CreateOptions) error {
	return ps.store.Problems().Create(ctx, prob, opt)
}

func (ps *probSrv) Update(ctx context.Context, prob *model.Problem, opt model.UpdateOptions) error {
	return ps.store.Problems().Update(ctx, prob, opt)
}

func (ps *probSrv) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Problem, error) {
	return ps.store.Problems().GetById(ctx, id, opt)
}

func (ps *probSrv) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ps.store.Problems().DeleteById(ctx, id, opt)
}

func (ps *probSrv) List(ctx context.Context, opt model.ListOptions) (*model.ProblemList, error) {
	return ps.store.Problems().List(ctx, opt)
}
