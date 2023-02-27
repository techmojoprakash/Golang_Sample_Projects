package cache

import (
	"context"
	"encoding/json"
	"redispoc/cmd/model"

	"github.com/go-redis/redis/v8"

	"fmt"
	"time"
)

var ctx = context.Background()

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		Host:    host,
		Db:      db,
		Expires: exp,
	}
}

type RedisCache struct {
	Host    string
	Db      int
	Expires time.Duration
}

func (c *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: "",
		DB:       c.Db,
	})
}

func (c *RedisCache) Set(key string, st *model.Student) error {
	client := c.getClient()

	// serialize Student object to JSON
	json, err := json.Marshal(st)
	if err != nil {
		fmt.Println("Unable marshal json ", err)
		return err
	}
	setErr := client.Set(ctx, key, json, c.Expires*time.Second).Err()
	if err != nil {
		fmt.Println("Unable to set the key into Redis Cache: setErr ", setErr)
		return err
	}
	return err
}

func (c *RedisCache) Get(key string) (*model.Student, error) {
	client := c.getClient()
	// fmt.Println("Data... ", cache, client)
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist in Redis Cache: err ", err)
		return nil, err
	}

	// serialize JSON to Student object
	stobj := &model.Student{}
	err = json.Unmarshal([]byte(val), stobj)
	if err != nil {
		fmt.Println("Unable Unmarshal stobj ", err)
		panic(err)
	}

	return stobj, nil
}

func (c *RedisCache) Del(key string) error {
	client := c.getClient()

	// checking key in cache
	_, err := c.Get(key)
	if err != nil {
		return err
	}

	val, delErr := client.Del(ctx, key).Result()
	if delErr != nil {
		fmt.Println("Unable to delete the key into Redis Cache: delErr ", delErr)
		return delErr
	}
	fmt.Println("Cache : Deleted Successfully", val)
	return nil
}
