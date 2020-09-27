package app

import (
	"kumparan-sbe-skilltest/config"
	newsController "kumparan-sbe-skilltest/internal/controller/http/news"
	newsLibrary "kumparan-sbe-skilltest/internal/library/news"
	newsRepository "kumparan-sbe-skilltest/internal/repository/news"
	newsRoute "kumparan-sbe-skilltest/routes/news"

	"github.com/labstack/echo"
)

// InitHTTPServer : Initialize HTTP server
func InitHTTPServer(config *config.Config, e *echo.Echo) {
	mysqlConn := ConnectMySQL(config)
	elasticSearchConn := ConnectElasticSearch(config)
	nsqPublisherConn := ConnectNSQPublisher(config)
	redisClient := ConnectRedis(config)

	// Repositories
	newsRepository := newsRepository.NewNewsRepository(nsqPublisherConn, mysqlConn, elasticSearchConn, redisClient)

	// Libraries
	newsLibrary := newsLibrary.NewNewsLibrary(newsRepository)

	// Controllers
	newsController := newsController.NewNewsController(newsLibrary)

	// Routes
	newsRoute.NewsRoutes(e, newsController)
}
