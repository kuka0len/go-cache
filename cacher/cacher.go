package cacher

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var ctx = context.TODO()

// Cacher -
type Cacher struct {
	prefix string
	ttl    time.Duration
}

// New -
func New(prefix string, ttl time.Duration) *Cacher {
	cacher := Cacher{prefix: prefix + ":", ttl: ttl}
	return &cacher
}

// Set -
func (cacher Cacher) Set(key string, value string) (string, error) {
	return client.Set(ctx, cacher.prefix+key, value, cacher.ttl).Result()
}

// Get -
func (cacher Cacher) Get(key string) (string, error) {
	return client.Get(ctx, cacher.prefix+key).Result()
}
