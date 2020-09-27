package app

import (
	"context"
	config "kumparan-sbe-skilltest/config"
	"kumparan-sbe-skilltest/helper"
	"log"

	// Import MySQL Driver
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nsqio/go-nsq"
	"github.com/olivere/elastic/v7"
)

// NSQPublisher : Wrapper for NSQ Publisher

// ConnectElasticSearch : Function for establishing connection to ElasticSearch
func ConnectElasticSearch(cfg *config.Config) config.ElasticConn {
	client, err := elastic.NewClient(elastic.SetURL("http://" + cfg.ElasticSearch.Host + ":" + cfg.ElasticSearch.Port))
	if err != nil {
		helper.WriteToLogFile("Failed to open ElasticSearch connection, ERROR :", err.Error())
		log.Fatal("Failed to open ElasticSearch connection, ERROR :", err.Error())
	}

	exists, err := client.IndexExists(cfg.ElasticSearch.Index).Do(context.Background())
	if err != nil {
		helper.WriteToLogFile("Failed to check ElasticSearch index, ERROR :", err.Error())
		log.Fatal("Failed to check ElasticSearch index, ERROR :", err.Error())
	}
	if !exists {
		mappings := `{
						"mappings": {
							"properties": {
								"news_id": { "type": "long" }
							}
						}
					}`
		_, err := client.CreateIndex(cfg.ElasticSearch.Index).Body(mappings).Do(context.Background())
		if err != nil {
			helper.WriteToLogFile("Failed to create ElasticSearch index, ERROR :", err.Error())
			log.Fatal("Failed to create ElasticSearch index, ERROR :", err.Error())
		}
	}

	return config.ElasticConn{
		ElasticConn: client,
		Index:       cfg.ElasticSearch.Index,
	}
}

// ConnectMySQL : Function for establishing connection to MySQL
func ConnectMySQL(config *config.Config) *sqlx.DB {
	db, err := sqlx.Open("mysql", config.MySQL.Username+":"+config.MySQL.Password+"@tcp("+config.MySQL.Host+":"+config.MySQL.Port+")/"+config.MySQL.Database)
	if err != nil {
		helper.WriteToLogFile("Failed to open MySQL connection, ERROR :", err.Error())
		log.Fatal("Failed to open MySQL connection, ERROR :", err.Error())
	}

	return db
}

// ConnectNSQPublisher : Function for establishing connection to NSQ Publisher
func ConnectNSQPublisher(cfg *config.Config) config.NSQPublisher {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(cfg.NSQPublisher.Host+":"+cfg.NSQPublisher.Port, nsqConfig)
	if err != nil {
		helper.WriteToLogFile("Failed to open new NSQ Publisher, ERROR :", err.Error())
		log.Fatal("Failed to open new NSQ Publisher, ERROR :", err.Error())
	}

	return config.NSQPublisher{
		Producer: producer,
		Topic:    cfg.NSQPublisher.Topic,
	}
}

// ConnectNSQSubscriber : Function for establishing connection to NSQ Subscriber
func ConnectNSQSubscriber(config *config.Config) *nsq.Consumer {
	nsqConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(config.NSQSubscriber.Topic, config.NSQSubscriber.Channel, nsqConfig)
	if err != nil {
		helper.WriteToLogFile("Failed to open new NSQ Subscriber, ERROR :", err.Error())
		log.Fatal("Failed to open new NSQ Subscriber, ERROR :", err.Error())
	}

	return consumer
}

// ConnectRedis : Function for establishing connection to Redis
func ConnectRedis(config *config.Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: "",
		DB:       0,
	})

	return redisClient
}
