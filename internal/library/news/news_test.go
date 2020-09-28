package news

import (
	"kumparan-sbe-skilltest/internal/interface/news/mocks"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"testing"
	"time"

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
	repoMock := new(mocks.Repository)

	repoMock.On("PublishNews", newsTestData).Return(nil)

	library := NewNewsLibrary(repoMock)
	err := library.PublishNews(newsTestData)

	assert.NoError(t, err)
}

func TestSaveNews(t *testing.T) {
	repoMock := new(mocks.Repository)

	newsTestData.Created = time.Now().Format("2006-01-02 15:04:05")

	newsElasicSearch := newsModel.NewsElasticSearch{
		NewsID:  0,
		Created: newsTestData.Created,
	}

	repoMock.On("SaveNews", newsTestData).Return(0, nil)
	repoMock.On("SaveNewsID", newsElasicSearch).Return(nil)

	library := NewNewsLibrary(repoMock)
	err := library.SaveNews(newsTestData)

	assert.NoError(t, err)
}

func TestGetNews(t *testing.T) {
	repoMock := new(mocks.Repository)

	var newsArray = make([]newsModel.News, 0)
	newsArray = append(newsArray, newsTestData)

	repoMock.On("GetNews", mock.AnythingOfType("int")).Return(newsArray, nil)

	library := NewNewsLibrary(repoMock)
	result, err := library.GetNews(1)

	assert.NoError(t, err)
	assert.Len(t, result, len(newsArray))
}
