package store

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
)

type ProbStore interface {
	Create(ctx context.Context, prob *model.Problem, opt model.CreateOptions) error
	Update(ctx context.Context, prob *model.Problem, opt model.UpdateOptions) error
	GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Problem, error)
	DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error
	List(ctx context.Context, opt model.ListOptions) (*model.ProblemList, error)
}
