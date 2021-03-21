package consumer

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	eventt "github.com/zehiraa/golang_test_project/infra/event"
	"github.com/zehiraa/golang_test_project/pkg/config"
	kafkaMocks "github.com/zehiraa/golang_test_project/pkg/queue/kafka/mocks"
	"github.com/zehiraa/golang_test_project/pkg/util"
	"time"
)

var _ = Describe("When first consumer", func() {
	var (
		message *kafka.Message
		err     error
	)

	Describe("And kafka message is valid", func() {
		var (
			event             eventt.TestEvent
			consumerContainer *kafkaMocks.ConsumerContainer
			producerContainer *kafkaMocks.ProducerContainer
		)
		util.BeforeAll(func() {
			var timeout time.Duration
			timeout = -1

			consumerContainer = new(kafkaMocks.ConsumerContainer)
			producerContainer = new(kafkaMocks.ProducerContainer)

			event = eventt.TestEvent{
				Name:    "test",
				SurName: "test",
			}
			jsonEvent, _ := json.Marshal(event)
			message = &kafka.Message{
				TopicPartition: kafka.TopicPartition{},
				Value:          jsonEvent,
				Key:            nil,
				Timestamp:      time.Time{},
				TimestampType:  0,
				Opaque:         nil,
				Headers:        nil,
			}

			consumerContainer.On("ReadMessage", timeout).Return(&message, nil).Maybe()
			consumerContainer.On("CommitMessage", message).Return([]kafka.TopicPartition{}, nil).Maybe()
			consumerContainer.On("Subscribe", "", mock.AnythingOfType("kafka.RebalanceCb")).Return(nil).Once()
			consumerContainer.On("Close").Return(nil).Once()

			producerContainer.On("ProduceAsync", message.Value, message.Headers, config.ProduceTopicName()).Return(nil).Once()

			firstConsumer := InitFirstConsumer(consumerContainer, producerContainer, "")
			err = firstConsumer.ProcessMessage(message)

		})

		It("should not return err", func() {
			Expect(err).Should(BeNil())
		})

		It("should call consumer commit method", func() {
			Expect(util.IsContain(consumerContainer.Calls, "CommitMessage")).Should(BeTrue())
		})
	})

	Describe("And kafka message is not valid", func() {
		var (
			consumerContainer *kafkaMocks.ConsumerContainer
			producerContainer *kafkaMocks.ProducerContainer
		)
		util.BeforeAll(func() {
			consumerContainer = new(kafkaMocks.ConsumerContainer)
			producerContainer = new(kafkaMocks.ProducerContainer)

			message = &kafka.Message{
				TopicPartition: kafka.TopicPartition{},
				Value:          nil,
				Key:            nil,
				Timestamp:      time.Time{},
				TimestampType:  1,
				Opaque:         nil,
				Headers:        nil,
			}

			consumerContainer.On("Subscribe", "", mock.AnythingOfType("kafka.RebalanceCb")).Return(nil).Once()
			consumerContainer.On("Close").Return(nil).Once()
			consumerContainer.On("CommitMessage", mock.Anything).Return([]kafka.TopicPartition{}, nil).Maybe()

			firstConsumer := InitFirstConsumer(consumerContainer, producerContainer, "")
			err = firstConsumer.ProcessMessage(message)

		})
		It("should return unmarshall err", func() {
			Expect(err).ShouldNot(BeNil())
		})

		It("should not call producer", func() {
			Expect(producerContainer.Calls).Should(HaveLen(0))
		})

		It("should call consumer commit method", func() {
			Expect(util.IsContain(consumerContainer.Calls, "CommitMessage")).Should(BeTrue())
		})
	})
})
