package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/ividernvi/algohub-forum/internal/model"
)

type tokenStore struct {
	rdb *redis.Client
}

func (ts *tokenStore) Create(ctx context.Context, token model.Token, opt model.CreateOptions) error {
	key := token.Token

	jsonToken, _ := json.Marshal(token)

	err := ts.rdb.SetNX(ctx, key, jsonToken, 2*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to store token in Redis: %v", err)
	}

	return nil
}

func (ts *tokenStore) Get(ctx context.Context, userid string, opt model.GetOptions) (*model.Token, error) {
	tokenstring, err := ts.rdb.Get(ctx, userid).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("token not found for user %d", userid)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get token from Redis: %v", err)
	}

	var token model.Token
	if err = json.Unmarshal([]byte(tokenstring), &token); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %s", err)
	}

	return &token, nil
}

func (ts *tokenStore) Delete(ctx context.Context, userid string, opt model.DeleteOptions) error {
	key := fmt.Sprintf("token:user:%d", userid)
	return ts.rdb.Del(ctx, key).Err()
}

func newTokens(ds *datastore) *tokenStore {
	return &tokenStore{
		rdb: ds.rdb,
	}
}
