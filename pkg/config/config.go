package config

import "os"

var configMap = make(map[string]interface{})

func Env() string {
	return readStringFromEnv("ENV", "default")
}

func Port() string {
	return readStringFromEnv("PORT", "8080")
}

func AppName() string {
	return readStringFromEnv("APP_NAME", "go-kafka-example-with-test")
}

func KafkaBroker() string {
	return readStringFromEnv("KAFKA_BROKER", "localhost:9092")
}

func KafkaGroupID() string {
	return readStringFromEnv("KAFKA_GROUP_ID", "go-kafka-example-with-test")
}

func KafkaAutoOffsetReset() string {
	return readStringFromEnv("KAFKA_AUTO_OFFSET_RESET", "latest")
}

func TestEventTopicName() string {
	return readStringFromEnv("CONSUMER_TEST_EVENT", "test-kafka-topic")
}
func ProduceTopicName() string {
	return readStringFromEnv("PRODUCE_TEST_EVENT", "produce-test-kafka-topic")
}

func readStringFromEnv(key string, defaultValue string) string {
	var value string

	if mapValue, exists := configMap[key]; exists {
		value = mapValue.(string)
	} else {
		value = defaultValue
		if envValue, exists := os.LookupEnv(key); exists {
			value = envValue
		}
		configMap[key] = value
	}

	return value
}