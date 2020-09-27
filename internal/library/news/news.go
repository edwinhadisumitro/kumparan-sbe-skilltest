package news

import (
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"time"
)

type newsLibrary struct {
	newsRepository newsInterface.Repository
}

// NewNewsLibrary : Establish new library for News
func NewNewsLibrary(newsRepository newsInterface.Repository) newsInterface.Library {
	return &newsLibrary{
		newsRepository: newsRepository,
	}
}

func (library *newsLibrary) PublishNews(news newsModel.News) error {
	var err error

	err = library.newsRepository.PublishNews(news)
	if err != nil {
		return err
	}

	return nil
}

func (library *newsLibrary) SaveNews(news newsModel.News) error {
	var err error

	created := time.Now().Format("2006-01-02 15:04:05")
	news.Created = created

	id, err := library.newsRepository.SaveNews(news)
	if err != nil {
		return err
	}

	err = library.newsRepository.SaveNewsID(newsModel.NewsElasticSearch{
		NewsID:  id,
		Created: created,
	})
	if err != nil {
		return err
	}

	return nil
}

func (library *newsLibrary) GetNews(page int) ([]newsModel.News, error) {
	var err error
	var news []newsModel.News

	news, err = library.newsRepository.GetNews(page)
	if err != nil {
		return news, err
	}

	return news, nil
}
