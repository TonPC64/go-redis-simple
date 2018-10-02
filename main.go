package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisSimpleClient struct {
	*redis.Client
}

func (r RedisSimpleClient) PingXTime(time int) {
	for i := 0; i < time; i++ {
		res, err := r.Ping().Result()
		fmt.Println(res, err)
	}
}

type RedisClusterClient struct {
	*redis.ClusterClient
}

func (r RedisClusterClient) PingXTime(time int) {
	for i := 0; i < time; i++ {
		res, err := r.Ping().Result()
		fmt.Println(res, err)
	}
}

type RedisClient interface {
	Ping() *redis.StatusCmd
	PingXTime(time int)
}

func NewRedisClient(isCluster bool) RedisClient {
	if isCluster {
		hosts := []string{""}
		password := ""
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    hosts,
			Password: password,
		})

		return RedisClient(RedisClusterClient{
			client,
		})
	}

	client := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	return RedisClient(RedisSimpleClient{
		client,
	})
}

func main() {
	NewRedisClient(true).PingXTime(5)
}
