package config

import (
	"fmt"
	"kumparan-sbe-skilltest/helper"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort       string
	MySQL         MySQL
	ElasticSearch ElasticSearch
	NSQPublisher  NSQ
	NSQSubscriber NSQ
	Redis         Redis
}

type MySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ElasticSearch struct {
	Host  string
	Port  string
	Index string
	Type  string
}

type NSQ struct {
	Host    string
	Port    string
	Topic   string
	Channel string
}

type Redis struct {
	Host string
	Port string
}

type NSQPublisher struct {
	Producer *nsq.Producer
	Topic    string
}

type ElasticConn struct {
	ElasticConn *elastic.Client
	Index       string
}

func ReadConfig() (*Config, error) {
	config := Config{}

	environment := strings.ToLower(os.Getenv("ENVIRONMENT"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// More option of config path can be added here
	viper.AddConfigPath(fmt.Sprintf("%s/config/%s/", *helper.ProjectFolder, environment))

	// Get the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Convert into struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
