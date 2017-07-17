package redispubsub

import (
	"github.com/DudeWhoCode/kulay/backend"
	. "github.com/DudeWhoCode/kulay/logger"
)

func Get(host string, port string, pass string, db int, channel string, rec chan<- string) {
	client := backend.NewRedisSession(host, port, pass, db)
	pubsub := client.Subscribe(channel)
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			Log.Error("Unable to receive message from redis pubsub channel\n", err)
		}
		if msg.Payload == "$^KILL^$" {
			break
		}
		rec <- msg.Payload
	}
}
