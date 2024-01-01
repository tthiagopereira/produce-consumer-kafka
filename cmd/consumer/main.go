package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "192.168.64.4:9092",
		"client.id":         "goapp-consumer", //para saber o quem estar consumindo o topico
		"group.id":          "goapp-group",    //

	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("Error consumer: ", err.Error())
	}

	topics := []string{
		"sms",
	}

	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}

}
