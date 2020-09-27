package routes

import (
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"

	"github.com/labstack/echo"
)

// NewsRoutes : Routes for news API
func NewsRoutes(e *echo.Echo, newsController newsInterface.HTTPController) {
	e.POST("/news", newsController.PublishNews)
	e.GET("/news", newsController.GetNews)
}
