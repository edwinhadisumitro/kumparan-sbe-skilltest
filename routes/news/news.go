package routes

import (
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"

	"github.com/labstack/echo"
)

// NewsRoutes : Routes for news API
func NewsRoutes(e *echo.Echo, newsController newsInterface.HTTPController) {
	newsGroup := e.Group("/news")

	newsGroup.POST("/publish", newsController.PublishNews)
	newsGroup.GET("/", newsController.GetNews)
}
