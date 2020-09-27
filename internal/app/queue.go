package app

import (
	config "kumparan-sbe-skilltest/config"
	newsController "kumparan-sbe-skilltest/internal/controller/queue/news"
	newsLibrary "kumparan-sbe-skilltest/internal/library/news"
	newsRepository "kumparan-sbe-skilltest/internal/repository/news"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// InitNSQSubscriber : Initialize NSQ Subscriber
func InitNSQSubscriber(config *config.Config) {
	mysqlConn := ConnectMySQL(config)
	elasticSearchConn := ConnectElasticSearch(config)
	nsqPublisherConn := ConnectNSQPublisher(config)
	nsqSubscriberConn := ConnectNSQSubscriber(config)
	redisClient := ConnectRedis(config)

	// Repositories
	newsRepository := newsRepository.NewNewsRepository(nsqPublisherConn, mysqlConn, elasticSearchConn, redisClient)

	// Libraries
	newsLibrary := newsLibrary.NewNewsLibrary(newsRepository)

	// NSQ Handler
	nsqSubscriberConn.AddHandler(newsController.NewNSQHandler(newsLibrary))

	// Must setup NSQ lookup after setting handler
	err := nsqSubscriberConn.ConnectToNSQLookupd(config.NSQSubscriber.Host + ":" + config.NSQSubscriber.Port)
	if err != nil {
		log.Fatal("Failed to connect to NSQ Lookup, ERROR :", err.Error())
	}

	// Wait for exit signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	nsqSubscriberConn.Stop()
}
