package redisq

import (
	"github.com/DudeWhoCode/kulay/backend"
	. "github.com/DudeWhoCode/kulay/logger"
	"time"
)


func Put(host string, port string, pass string, db int, queue string, rec <-chan string) {
	client := backend.NewRedisSession(host, port, pass, db)
	for msg := range rec {
		if err := client.RPush(queue, msg).Err(); err != nil {
			Log.Warn(err)
		}
	}
}


func Get(host string, port string, pass string, db int, queue string, rec chan<- string) {
	client := backend.NewRedisSession(host, port, pass, db)
	// use `client.BLPop(0, "queue")` for infinite waiting time
	for {
		result, err := client.BLPop(1*time.Second, queue).Result()
		if err != nil {
			Log.Warn(err)
		}
		if result == nil {
			Log.Info("Received all messages from redis queue")
			break
		}
		rec <- result[1]
	}
}