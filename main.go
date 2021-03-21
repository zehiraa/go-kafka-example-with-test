package main

import (
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/zehiraa/golang_test_project/appserver"
	"github.com/zehiraa/golang_test_project/bootstraper"
	"github.com/zehiraa/golang_test_project/infra/consumer"
	"github.com/zehiraa/golang_test_project/pkg/config"
	"github.com/zehiraa/golang_test_project/pkg/queue/kafka"
	"os"
)

var (
	kafkaConsumer kafka.ConsumerContainer
	kafkaProducer kafka.ProducerContainer
	logger        log.Logger
)

func init() {
	kafkaConsumer = bootstraper.LoadConsumer()
	kafkaProducer = bootstraper.LoadProducer()
	logger = log.Logger{
		Out:          os.Stdout,
		Hooks:        nil,
		Formatter:    &log.JSONFormatter{},
		ReportCaller: false,
		Level:        4,
		ExitFunc:     nil,
	}
}

func main() {

	logger.Info("start main func ...")

	go consumer.InitFirstConsumer(kafkaConsumer, kafkaProducer, config.TestEventTopicName()).Consume()

	router := chi.NewRouter()
	startServer(router, &logger)
}

func startServer(router *chi.Mux, logger *log.Logger) {
	server := &appserver.Server{
		Router: router,
		Logger: logger,
	}
	server.HealthCheck()
	server.Start()
}
