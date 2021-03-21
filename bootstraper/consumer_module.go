package bootstraper

import (
	"github.com/zehiraa/golang_test_project/pkg/config"
	"github.com/zehiraa/golang_test_project/pkg/queue/kafka"
)

func LoadConsumer() kafka.ConsumerContainer{

	testConsumer, err := kafka.NewConsumer(&kafka.ConsumerConfig{
		Broker:          config.KafkaBroker(),
		GroupID:         config.KafkaGroupID(),
		AutoOffsetReset: config.KafkaAutoOffsetReset(),
	})

	if err != nil {
		println(err.Error())
		panic("cannot start test kafka")
	}
	return testConsumer
}