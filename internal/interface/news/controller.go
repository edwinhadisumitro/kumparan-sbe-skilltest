package news

import "github.com/labstack/echo"

// HTTPController : HTTP Controller Interface for News
type HTTPController interface {
	PublishNews(c echo.Context) error
}
