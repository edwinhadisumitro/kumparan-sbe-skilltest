package news

import newsModel "kumparan-sbe-skilltest/internal/model/news"

// Library : Library Interface for News
type Library interface {
	PublishNews(news newsModel.News) error
	SaveNews(news newsModel.News) error
}
