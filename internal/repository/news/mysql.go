package news

import (
	newsModel "kumparan-sbe-skilltest/internal/model/news"
)

func (repo *newsRepository) SaveNews(news newsModel.News) (int, error) {
	var err error
	var lastInsertID int

	query := `insert into news (author, body, created) values (?, ?, ?);`
	res, err := repo.mysqlConn.Exec(query, news.Author, news.Body, news.Created)
	if err != nil {
		return lastInsertID, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return lastInsertID, err
	}

	lastInsertID = int(id)

	return lastInsertID, nil
}

func (repo *newsRepository) GetNewsDetail(newsElastic newsModel.NewsElasticSearch) (newsModel.News, error) {
	news := newsModel.NewsSQL{}
	query := `select id, author, body, created from news where id=?`
	err := repo.mysqlConn.Get(&news, query, newsElastic.NewsID)
	if err != nil {
		return news.ConvertToJSON(), err
	}

	return news.ConvertToJSON(), nil
}
