package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type ProducerContainer interface {
	ProduceAsync(message []byte, headers []kafka.Header, topic string) error
}
type producerContainer struct {
	Producer       *kafka.Producer
	configurations *ConsumerConfig
	isConnected    bool
}

func NewProducer(configurations *ConsumerConfig) (ProducerContainer,error){

	var err error
	producerCont := &producerContainer{configurations: configurations}

	producerCont.Producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": configurations.Broker,
		"acks":              "all"})

	if err != nil {
		log.Error(fmt.Sprintf("an error occurred while connecting kafka, err: %v", err))
	}

	producerCont.isConnected = true
	return producerCont, nil
}


func (c *producerContainer) ProduceAsync(message []byte, headers []kafka.Header, topic string) error {
	deliveryChan := make(chan kafka.Event)
	var err error
	defer close(deliveryChan)

	if err = c.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
		Headers:        headers,
	},
		deliveryChan,
	); err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Error(fmt.Sprintf("Delivery failed: %v\n %v", m.TopicPartition.Error, topic))
		err = m.TopicPartition.Error
	} else {
		log.Info(fmt.Sprintf("Delivered message to topic %s [%d] at offset %d\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
	}

	return err
}
