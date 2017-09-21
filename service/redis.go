/**
 * message queue
 */
package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
)

var RedisKey = "go_git_webhooks_key"

func RedisClient(addr string, password string, db int) (r *redis.Client, err error) {
	rClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	pong, err := rClient.Ping().Result()
	fmt.Println(pong, reflect.TypeOf(pong), err, reflect.TypeOf(err))
	// Output: PONG <nil>
	return rClient, err
}
