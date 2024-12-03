package store

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
)

type TokenStore interface {
	Create(ctx context.Context, token model.Token, opt model.CreateOptions) error
	Get(ctx context.Context, userid string, opt model.GetOptions) (*model.Token, error)
	Delete(ctx context.Context, userid string, opt model.DeleteOptions) error
}
