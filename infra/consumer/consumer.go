package consumer

import (
	"encoding/json"
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	eventPackage "github.com/zehiraa/golang_test_project/infra/event"
	"github.com/zehiraa/golang_test_project/pkg/config"
	"github.com/zehiraa/golang_test_project/pkg/queue/kafka"
)

type firstConsumer struct {
	consumer kafka.ConsumerContainer
	producer kafka.ProducerContainer
	topic    string
}

type FirstConsumer interface {
	Consume()
	ProcessMessage(msg *confluentKafka.Message) error
}

func InitFirstConsumer(consumer kafka.ConsumerContainer, producer kafka.ProducerContainer, topic string) FirstConsumer {
	return &firstConsumer{
		consumer: consumer,
		producer: producer,
		topic:    topic,
	}
}

func (c *firstConsumer) Consume() {
	_ = c.consumer.Subscribe(c.topic, nil)
	defer c.consumer.Close()
	for {
		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			log.Error(fmt.Sprintf("could not read message, error: %v", err))
			continue
		}
		log.Info("retrieve message")

		err = c.ProcessMessage(msg)
		if err != nil {
			continue
		}
	}
}

func (c *firstConsumer) ProcessMessage(msg *confluentKafka.Message) error {

	var err error
	var event eventPackage.TestEvent

	err = json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Error(fmt.Sprintf("serialization error: %v", err))
		_, _ = c.consumer.CommitMessage(msg)
		return err
	}

	err = c.producer.ProduceAsync(msg.Value, msg.Headers, config.ProduceTopicName())
	if err == nil {
		_, _ = c.consumer.CommitMessage(msg)

		log.Info(fmt.Sprintf("process complete. event: %v", event))

		return nil
	}
	return err
}
