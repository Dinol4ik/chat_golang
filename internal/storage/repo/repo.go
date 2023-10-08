package repo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Repo struct {
	db *redis.Client
}

func NewConnect(db *redis.Client) Repo {
	return Repo{db}
}

func (r *Repo) AddMessage(ctx context.Context, message string, param string) error {
	err := r.db.RPush(ctx, "chat:"+param, message).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *Repo) TakeLastMessages(ctx context.Context, param string) ([]string, error) {
	e := r.db.Info(ctx)
	log.Println(e)
	appe := r.db.Set(ctx, "foo", "dsad", -1)
	_ = appe
	lastMessages := r.db.LLen(ctx, "chat:"+param).Val()
	result, err := r.db.LRange(ctx, "chat:"+param, lastMessages-50, lastMessages).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
