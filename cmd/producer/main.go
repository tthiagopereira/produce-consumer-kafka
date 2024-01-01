package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	produce := NewKafkaProduce()
	err := Publish("messagem - 4", "sms", produce, []byte("transferencia"), deliveryChan)
	if err != nil {
		return
	}
	go DeliveryReport(deliveryChan)

	log.Println("Send produce kafka success")
	produce.Flush(1000)
}

func NewKafkaProduce() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "192.168.64.4:9092",
		"delivery.timeout.ms": "0",    //Tempo maximo que espera para retorno
		"acks":                "all",  // 0, 1, all 0 -> nao precisa receber o retorno, 1 -> lider informa que ja chegou a messagem, all -> lider informa que ja chegou a menssagem depois que ja foi feita a replicacao para todas as instancias
		"enable.idempotence":  "true", // para ter a certeza que a mensagem foi enviada com sucesso
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Fatal("Error connecting to Kafka")
		log.Println(err.Error())
		panic(err)
	}
	log.Println("Instance apache kafka success")
	return p
}

func Publish(msg, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	log.Println("[DeliveryReport] - [] ")
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Println("Error ao  enviar")
				// fazer processo para reenvio ou salvar no banco de dados para envios futuros
			} else {
				log.Println("Menssagem enviada: ", ev.TopicPartition)
				// informar que foi enviado
			}
		}
	}
}
