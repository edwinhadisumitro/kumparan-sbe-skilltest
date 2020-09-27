package news

import (
	"encoding/json"
	"kumparan-sbe-skilltest/internal/model/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"strconv"
	"time"
)

func (repo *newsRepository) GetNewsFromCache(newsElastic newsModel.NewsElasticSearch) (news.News, error) {
	var news newsModel.News
	result, err := repo.redisClient.Get("NEWS:" + strconv.Itoa(newsElastic.NewsID)).Result()
	if err != nil {
		return news, err
	}

	err = json.Unmarshal([]byte(result), &news)
	if err != nil {
		return news, err
	}

	return news, nil
}

func (repo *newsRepository) SaveNewsToCache(news newsModel.News) error {
	data, err := json.Marshal(news)
	if err != nil {
		return err
	}

	repo.redisClient.Set("NEWS:"+strconv.Itoa(news.ID), string(data), 60*time.Second)

	return nil
}
