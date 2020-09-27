package news

import (
	"kumparan-sbe-skilltest/internal/model/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
)

// Repository : Repository Interface for News
type Repository interface {
	PublishNews(news newsModel.News) error
	SaveNews(news newsModel.News) (int, error)
	SaveNewsID(news newsModel.NewsElasticSearch) error
	GetNews(page int) ([]newsModel.News, error)
	GetNewsFromCache(newsElastic newsModel.NewsElasticSearch) (news.News, error)
	SaveNewsToCache(news newsModel.News) error
	GetNewsDetail(newsElastic newsModel.NewsElasticSearch) (newsModel.News, error)
}
