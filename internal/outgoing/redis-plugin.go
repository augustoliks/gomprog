package outgoing

import (
	"context"
	"encoding/json"
	"time"

	"github.com/augustoliks/gomprog/internal/service"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient *redis.Client

type RedisPlugin struct {
	URL      string
	Password string
}

func (redisPlugin RedisPlugin) OnInit() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisPlugin.URL,
		Password: redisPlugin.Password,
	})

	err := redisClient.Ping(context.Background()).Err()

	if err != nil {
		time.Sleep(3 * time.Second)
		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			panic(err)
		}
	}

}

func (redisPlugin RedisPlugin) OnSend(log service.GELFLogFormat) {
	jsonLog, err := json.Marshal(log)

	if err != nil {
		panic(err)
	}

	err = redisClient.Publish(ctx, log.AppName, string(jsonLog)).Err()

	if err != nil {
		panic(err)
	}

}
