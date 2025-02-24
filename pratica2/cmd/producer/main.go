package main

// #cgo LDFLAGS: -L/usr/lib -lrdkafka
import "C"
import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("Hello, kafka!", "teste", producer, []byte("transferencia879"), deliveryChan)
	go DeliveryReport(deliveryChan)
	// //Isso aqui é sincrono, entao o codigo so continua quando receber uma msg pelo canal
	// e := <-deliveryChan
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	log.Println("Erro ao enviar a mensagem", msg.TopicPartition.Error)
	// } else {
	// 	log.Println("Mensagem enviada", msg.TopicPartition)
	// }

	fmt.Println("Mensagem enviada")
	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true", //Por padrao ele é falso
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Println("Erro ao enviar a mensagem", ev.TopicPartition.Error)
			} else {
				//Aqui se pode ter notacoes de banco de dados para gravar e saber que a mensagem foi enviada
				log.Println("Mensagem enviada", ev.TopicPartition)
			}
		}
	}
}
