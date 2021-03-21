package bootstraper

import (
	"github.com/zehiraa/golang_test_project/pkg/config"
	"github.com/zehiraa/golang_test_project/pkg/queue/kafka"
)

func LoadProducer() kafka.ProducerContainer {

	producer, err := kafka.NewProducer(
		&kafka.ConsumerConfig{Broker: config.KafkaBroker()})

	if err != nil {
		println(err.Error())
		panic("cannot create producer")
	}

	return producer
}
