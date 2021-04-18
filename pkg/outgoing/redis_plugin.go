package outgoing

import (
	"context"
	"encoding/json"
	"time"

	"github.com/augustoliks/gomprog/pkg/service"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient *redis.Client

type RedisPlugin struct {
	URL      string
	Password string
}

func (redisPlugin RedisPlugin) OnInit() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisPlugin.URL,
		Password: redisPlugin.Password,
	})

	err := redisClient.Ping(context.Background()).Err()

	if err != nil {
		time.Sleep(5 * time.Second)
		err = redisClient.Ping(context.Background()).Err()
	}

	return err
}

func (redisPlugin RedisPlugin) OnSend(log service.GELFLogFormat) error {
	jsonLog, err := json.Marshal(log)
	err = redisClient.Publish(ctx, log.AppName, string(jsonLog)).Err()
	return err
}
