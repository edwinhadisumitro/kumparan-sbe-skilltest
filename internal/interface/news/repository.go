package news

import (
	newsModel "kumparan-sbe-skilltest/internal/model/news"
)

// Repository : Repository Interface for News
type Repository interface {
	PublishNews(news newsModel.News) error
	SaveNews(news newsModel.News) (int, error)
	SaveNewsID(id int) error
}
