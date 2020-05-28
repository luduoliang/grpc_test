package redis

import (
	"grpc_test/config"
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"strconv"
)

var prefix string

var redis *radix.Pool

func init() {
	prefix = config.Get("REDIS_PREFIX") + ":"
	redis, _ = radix.NewPool("tcp", config.Get("REDIS_HOST")+":"+config.Get("REDIS_PORT"), 100, nil)
}

//Subscribe Subscribe to a channel
func Subscribe(channel string) chan radix.PubSubMessage {
	conn, _ := radix.Dial("tcp", config.Get("REDIS_HOST")+":"+config.Get("REDIS_PORT"))
	ps := radix.PubSub(conn)
	msgSignal := make(chan radix.PubSubMessage)
	_ = ps.Subscribe(msgSignal, channel)
	return msgSignal
}

//Publish publish data to a channel
func Publish(channel, data string) {
	_ = redis.Do(radix.Cmd(nil, "PUBLISH", channel, data))
}

//Get get a value
func Get(key string) string {
	var r string
	err := redis.Do(radix.Cmd(&r, "GET", prefix+key))
	if err != nil {
		return ""
	}
	return r
}

//GetInt get a value as int
func GetInt(key string) int {
	var r string
	err := redis.Do(radix.Cmd(&r, "GET", prefix+key))
	if err != nil {
		return 0
	}
	intR, _ := strconv.Atoi(r)
	return int(intR)
}

//Set set a value
func Set(key string, val string, ttl int) {
	_ = redis.Do(radix.Cmd(nil, "SET", prefix+key, val))
	if ttl > 0 {
		_ = redis.Do(radix.Cmd(nil, "EXPIRE", prefix+key, fmt.Sprintf("%d", ttl)))
	}
}

//Expire a key
func Expire(key string, ttl int) {
	_ = redis.Do(radix.Cmd(nil, "EXPIRE", prefix+key, fmt.Sprintf("%d", ttl)))
}

//HSet Hset a value
func HSet(key string, val string, val2 string) {
	_ = redis.Do(radix.Cmd(nil, "HSET", prefix+key, val, val2))
}

//HIncrBy HINCRBY
func HIncrBy(key string, field string, increment string) {
	_ = redis.Do(radix.Cmd(nil, "HINCRBY", prefix+key, field, increment))
}

//HDel Hdel a value
func HDel(key string, val string) {
	_ = redis.Do(radix.Cmd(nil, "HDEL", prefix+key, val))
}

//HGet Hget a value
func HGet(key string, val string) string {
	var r string
	err := redis.Do(radix.Cmd(&r, "HGET", prefix+key, val))
	if err != nil {
		return ""
	}
	return r
}

//HGetAll Hgetall a value
func HGetAll(key string) map[string]string {
	var r map[string]string
	err := redis.Do(radix.Cmd(&r, "HGETALL", prefix+key))
	if err != nil {
		return r
	}
	return r
}

//HLen Get the size of hash
func HLen(key string) int {
	var r int
	err := redis.Do(radix.Cmd(&r, "HLEN", prefix+key))
	if err != nil {
		return 0
	}
	return r
}

//ZAdd zadd
func ZAdd(key string, score string, member string) {
	_ = redis.Do(radix.Cmd(nil, "ZADD", prefix+key, score, member))
}

//ZRem zrem
func ZRem(key string, member string) {
	_ = redis.Do(radix.Cmd(nil, "ZREM", prefix+key, member))
}

//ZScore zscore
func ZScore(key string, member string) int {
	var r int
	err := redis.Do(radix.Cmd(&r, "ZSCORE", prefix+key, member))
	if err != nil {
		return 0
	}
	return r
}

//ZRank zrank
func ZRank(key string, member string) int {
	var r int
	err := redis.Do(radix.Cmd(&r, "ZRANK", prefix+key, member))
	if err != nil {
		return 0
	}
	return r
}

//ZRevRank zrank
func ZRevRank(key string, member string) int {
	var r int
	err := redis.Do(radix.Cmd(&r, "ZREVRANK", prefix+key, member))
	if err != nil {
		return 0
	}
	return r
}

//ZCount zcount
func ZCount(key string) int {
	var r int
	err := redis.Do(radix.Cmd(&r, "ZCOUNT", prefix+key, "-inf", "+inf"))
	if err != nil {
		return 0
	}
	return r
}

//ZRevRange zrevrange
func ZRevRange(key string, startScore, endScore string) []string {
	var r []string
	err := redis.Do(radix.Cmd(&r, "ZREVRANGE", prefix+key, startScore, endScore))
	if err != nil {
		r = make([]string, 0)
	}
	return r
}

//ZRevRangeByScore zrevrangebyscore
func ZRevRangeByScore(key string, startScore, endScore, limit string) []string {
	var r []string
	err := redis.Do(radix.Cmd(&r, "ZREVRANGEBYSCORE", prefix+key, "("+startScore, endScore, "LIMIT", "0", limit))
	if err != nil {
		r = make([]string, 0)
	}
	return r
}
