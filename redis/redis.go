package redis

import "github.com/go-redis/redis"
import "context"

var (
	client *redis.Client
	ctx context.Context
)

func InitRedisClient(addr,passwd string) error {

	ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	return err
}

func GetKeys(key string)(string,error)  {

	res,err := client.Get(ctx,key).Result()

	return res,err
}