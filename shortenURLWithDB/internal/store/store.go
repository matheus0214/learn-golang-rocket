package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GeFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(db *redis.Client) Store {
	return store{rdb: db}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string

	for range 5 {
		code = genCode()

		if err := s.rdb.HGet(ctx, "encoded", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}

			return "", fmt.Errorf("failed to get code from encoded hashmap: %w", err)
		}
	}

	if err := s.rdb.HSet(ctx, "encoded", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in encoded hashmap: %w", err)
	}

	return code, nil
}

func (s store) GeFullURL(ctx context.Context, code string) (string, error) {
	_url, err := s.rdb.HGet(ctx, "encoded", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get code from encoded hashmap: %w", err)
	}

	return _url, nil
}
