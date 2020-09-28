package news

import (
	"encoding/json"
	"kumparan-sbe-skilltest/internal/interface/news/mocks"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var newsTestData = newsModel.News{
	ID:      1,
	Author:  "Foo Bar",
	Body:    "Kumparan SBE Skilltest",
	Created: time.Now().Format("2006-01-02 15:04:05"),
}

func TestPublishNews(t *testing.T) {
	libraryMock := new(mocks.Library)
	libraryMock.On("PublishNews", newsTestData).Return(nil)

	newsDataByte, _ := json.Marshal(newsTestData)

	e := echo.New()
	req, err := http.NewRequest("POST", "/news", strings.NewReader(string(newsDataByte)))
	req.Header.Add("Content-Type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/news")

	controller := NewNewsController(libraryMock)
	err = controller.PublishNews(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetNews(t *testing.T) {
	newsArray := make([]newsModel.News, 0)
	newsArray = append(newsArray, newsTestData)

	libraryMock := new(mocks.Library)
	libraryMock.On("GetNews", mock.AnythingOfType("int")).Return(newsArray, nil)

	e := echo.New()
	req, err := http.NewRequest("GET", "/news", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("page", "1")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/news")

	controller := NewNewsController(libraryMock)
	err = controller.GetNews(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
