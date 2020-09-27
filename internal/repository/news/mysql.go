package news

import (
	newsModel "kumparan-sbe-skilltest/internal/model/news"
)

func (repo *newsRepository) SaveNews(news newsModel.News) (int, error) {
	var err error
	var lastInsertID int

	query := `insert into news (author, body, created) values (?, ?, current_timestamp());`
	res, err := repo.mysqlConn.Exec(query, news.Author, news.Body)
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
