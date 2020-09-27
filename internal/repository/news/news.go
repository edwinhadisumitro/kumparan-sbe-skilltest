package news

import (
	"kumparan-sbe-skilltest/config"
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"

	"github.com/jmoiron/sqlx"
)

type newsRepository struct {
	nsqProducer config.NSQPublisher
	mysqlConn   *sqlx.DB
	elasticConn config.ElasticConn
}

// NewNewsRepository : Establish new repository for News
func NewNewsRepository(nsq config.NSQPublisher, mysqlConn *sqlx.DB, elasticConn config.ElasticConn) newsInterface.Repository {
	return &newsRepository{
		nsqProducer: nsq,
		mysqlConn:   mysqlConn,
		elasticConn: elasticConn,
	}
}
