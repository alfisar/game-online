package repository

import "time"

type RedisRepositoryContract interface {
	Insert(key string, value string, exp time.Duration) (err error)
	Get(key string) (data string, err error)
	Delete(key string) (err error)
}
