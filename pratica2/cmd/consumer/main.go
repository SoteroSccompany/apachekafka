package main

// #cgo LDFLAGS: -L/usr/lib -lrdkafka
import "C"
import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp2-group",
		"auto.offset.reset": "earliest", //Aqui ele vai pegar todas as msgs
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("Erro ao criar o consumer", err.Error())
	}
	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}

}
