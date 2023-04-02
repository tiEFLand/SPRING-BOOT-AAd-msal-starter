package utils

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisUtils struct {
	conn redis.Conn
}

func (r *RedisUtils) Connect() {
	url, _ := GetConfig().Get("redis.url")
	c, err := redis.DialURL(url)
	for err != nil {
		logger.Error(err)
		time.Sleep(time.Second * 2)
		if c == nil || c.Err() != nil {
			c, err = redis.DialURL(url)
		}
	}
	r.conn = c
}

func (r *RedisUtils) Close()