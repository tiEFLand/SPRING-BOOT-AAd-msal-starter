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
	url, _ := G