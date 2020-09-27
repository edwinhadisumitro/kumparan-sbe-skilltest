package news

import (
	"kumparan-sbe-skilltest/helper"
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nsqio/go-nsq"
)

type newsController struct {
	newsLibrary       newsInterface.Library
	nsqSubscriberConn *nsq.Consumer
}

// NewNewsController : Establish new controller for News
func NewNewsController(newsLibrary newsInterface.Library) newsInterface.HTTPController {
	return &newsController{
		newsLibrary: newsLibrary,
	}
}

func (controller *newsController) PublishNews(c echo.Context) error {
	var err error
	var news newsModel.News
	var response helper.HTTPResponse
	var responsePayload map[string]interface{}

	err = c.Bind(&news)
	if err != nil {
		response = helper.NewErrorResponse("Error parsing request body", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	err = controller.newsLibrary.PublishNews(news)
	if err != nil {
		response = helper.NewErrorResponse("Failed to publish news", err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response = helper.NewSuccessResponse("", responsePayload)
	return c.JSON(http.StatusOK, response)
}
