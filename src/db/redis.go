package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"golasc/src/appconf"
)

var ctx = context.Background()

func MakeClient() (*context.Context, *redis.Client) {
	return &ctx, redis.NewClient(&redis.Options{
		Addr:     appconf.Config.Redis.Host + ":" + appconf.Config.Redis.Port,
		Password: appconf.Config.Redis.Pass,
		DB:       appconf.Config.Redis.Db,
	})
}
