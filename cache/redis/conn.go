package redis

import "github.com/garyburd/redigo/redis"

var (
	pool *redis.Pool
	redisHost = "1.15.90.134:6379"
	redisPass = ""
)
