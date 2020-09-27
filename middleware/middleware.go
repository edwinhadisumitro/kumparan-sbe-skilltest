package middleware

import (
	"kumparan-sbe-skilltest/helper"
	"net/http"

	"github.com/labstack/echo"
)

// Middleware : Middleware struct wrapper
type Middleware struct {
	projectFolder string
}

// NewMiddleware : Initialize Echo's middleware
func NewMiddleware(projectFolder *string) *Middleware {
	return &Middleware{
		projectFolder: *projectFolder,
	}
}

// ValidateContentType : Validate ContentType to application/json if request method not GET
func (m *Middleware) ValidateContentType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		if c.Request().Method != "GET" {
			if contentType != "application/json" {
				response := helper.NewErrorResponse("Content-Type not supported", "ContentType must be application/json")

				return c.JSON(http.StatusBadRequest, response)
			}
		}

		return next(c)
	}
}
