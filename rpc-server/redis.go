package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	cli *redis.Client
}

func (c *RedisClient) InitClient(ctx context.Context, address, password string) error {
	r := redis.NewClient(&redis.Options{
		Addr: address,
		// no pw set
		Password: password,
		// use default DB
		DB: 0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		return err
	}

	c.cli = r
	return nil
}
