package confs

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func SetupRedisConnection() *redis.Client {
	RedisHost := os.Getenv("REDIS_HOST")
	RedisPort := os.Getenv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", RedisHost, RedisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	return rdb
}

func CloseRedisConnection(rdb *redis.Client) error {
	return rdb.Close()
}

func SetupRedisClusterConnection() *redis.ClusterClient {
	RedisAddrs := os.Getenv("CLU_REDIS_ADDRS")
	RedisPasswd := os.Getenv("CLU_REDIS_PASSWD")

	//[]string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

	addrs := strings.Split(RedisAddrs, ",")

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        addrs,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		Password:     RedisPasswd,
		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})
	// addr := fmt.Sprintf("%s:%s", RedisHost, RedisPort)

	// rdb := redis.NewClient(&redis.Options{
	//      Addr:         addr,
	//      DialTimeout:  10 * time.Second,
	//      ReadTimeout:  30 * time.Second,
	//      WriteTimeout: 30 * time.Second,
	//      PoolSize:     10,
	//      PoolTimeout:  30 * time.Second,
	// })

	cmd := rdb.Ping(context.Background())

	fmt.Println(cmd.Result())

	return rdb
}

func CloseRedisClusterConnection(rdb *redis.ClusterClient) error {
	return rdb.Close()
}
