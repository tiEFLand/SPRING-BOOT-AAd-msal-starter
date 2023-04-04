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

func (r *RedisUtils) Close() {
	if r.conn == nil {
		return
	}
	r.conn.Close()
}

//127.0.0.1:6379[1]> get "depth:OKEX:spot:XRP"
func (r *RedisUtils) SetDepth(key string, expireTime int, data interface{}) {
	if r.conn == nil {
		return
	}
	_, err := r.conn.Do("psetex", key, expireTime, data)
	if err != nil {
		logger.Error(err)
	}
}

//127.0.0.1:6379[1]> get "depth:OKEX:spot:XRP"
func (r *RedisUtils) GetDepth(key string) (string, error) {
	if r.conn == nil {
		return "", errors.New("redis disconnected")
	}
	json, getErr := redis.String(r.conn.Do("get", key))
	return json, getErr
}
