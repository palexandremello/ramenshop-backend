package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisPublisherEvent is a struct that represents a RedisPublisherEvent
type RedisPublisherEvent struct {
	client *redis.Client
}

// NewRedisPublisherEvent is a factory function that creates a new RedisPublisherEvent
func NewRedisPublisherEvent(addr string) *RedisPublisherEvent {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisPublisherEvent{client: client}
}

// Execute is a method that executes the RedisPublisherEvent
func (rp *RedisPublisherEvent) Execute(channel string, message string) error {
	return rp.client.Publish(context.Background(), channel, message).Err()
}
