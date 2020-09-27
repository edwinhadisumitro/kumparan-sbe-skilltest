package news

import (
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
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

	id, err := library.newsRepository.SaveNews(news)
	if err != nil {
		return err
	}

	err = library.newsRepository.SaveNewsID(id)
	if err != nil {
		return err
	}

	return nil
}
