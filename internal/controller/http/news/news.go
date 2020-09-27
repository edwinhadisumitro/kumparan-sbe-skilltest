package news

import (
	"kumparan-sbe-skilltest/helper"
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"net/http"
	"strconv"

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

func (controller *newsController) GetNews(c echo.Context) error {
	var err error
	var news []newsModel.News
	var page int
	var response helper.HTTPResponse

	page, err = strconv.Atoi(c.Request().Header.Get("page"))
	if err != nil {
		response = helper.NewErrorResponse("Failed to get page number", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	if page == 0 {
		response = helper.NewErrorResponse("Failed to validate page number", "Page number must be greater than 0")
		return c.JSON(http.StatusBadRequest, response)
	}

	news, err = controller.newsLibrary.GetNews(page)
	if err != nil {
		response = helper.NewErrorResponse("Failed to get News data", err.Error())
	}

	response = helper.NewSuccessResponse("news", news)
	return c.JSON(http.StatusOK, response)
}
