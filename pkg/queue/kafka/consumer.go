package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type ConsumerContainer interface {
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	Commit() ([]kafka.TopicPartition, error)
	Close() (err error)
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
	CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error)
}

type consumerContainer struct {
	Consumer *kafka.Consumer
}

type ConsumerConfig struct {
	Broker          string
	GroupID         string
	AutoOffsetReset string
}

func NewConsumer(config *ConsumerConfig) (ConsumerContainer, error) {
	var err error

	container := consumerContainer{}
	container.Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.Broker,
		"group.id":           config.GroupID,
		"auto.offset.reset":  config.AutoOffsetReset,
		"enable.auto.commit": "false",
	})

	return &container, err
}

func (c *consumerContainer) Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error {
	return c.Consumer.Subscribe(topic, rebalanceCb)
}

func (c *consumerContainer) Commit() ([]kafka.TopicPartition, error) {
	return c.Consumer.Commit()
}

func (c *consumerContainer) Close() (err error) {
	return c.Consumer.Close()
}
func (c *consumerContainer) ReadMessage(timeout time.Duration) (*kafka.Message, error) {
	return c.Consumer.ReadMessage(timeout)
}

func (c *consumerContainer) CommitMessage(m *kafka.Message) ([]kafka.TopicPartition, error) {
	return c.Consumer.CommitMessage(m)
}