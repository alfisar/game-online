package repository

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	conn *redis.Client
}

func NewRedisRepository(conn *redis.Client) *RedisRepository {
	return &RedisRepository{
		conn: conn,
	}
}

func (obj *RedisRepository) Insert(key string, value string, exp time.Duration) (err error) {
	result := obj.conn.Set(key, value, exp)
	if result.Err() != nil {
		err = result.Err()
		return
	}

	return
}

func (obj *RedisRepository) Get(key string) (data string, err error) {
	data, err = obj.conn.Get(key).Result()
	if err != nil {
		return
	}
	return
}

func (obj *RedisRepository) Delete(key string) (err error) {
	result := obj.conn.Del(key)
	if result.Err() != nil {
		err = result.Err()
		return
	}

	return
}
