package news

import (
	"kumparan-sbe-skilltest/config"
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type newsRepository struct {
	nsqProducer config.NSQPublisher
	mysqlConn   *sqlx.DB
	elasticConn config.ElasticConn
	redisClient *redis.Client
}

// NewNewsRepository : Establish new repository for News
func NewNewsRepository(nsq config.NSQPublisher, mysqlConn *sqlx.DB, elasticConn config.ElasticConn, redisClient *redis.Client) newsInterface.Repository {
	return &newsRepository{
		nsqProducer: nsq,
		mysqlConn:   mysqlConn,
		elasticConn: elasticConn,
		redisClient: redisClient,
	}
}
