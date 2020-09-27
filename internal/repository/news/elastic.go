package news

import (
	"context"
	"encoding/json"
	"kumparan-sbe-skilltest/helper"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func (repo *newsRepository) SaveNewsID(news newsModel.NewsElasticSearch) error {
	var err error

	_, err = repo.elasticConn.ElasticConn.Index().Index(repo.elasticConn.Index).BodyJson(news).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (repo *newsRepository) GetNews(page int) ([]newsModel.News, error) {
	var news []newsModel.News
	dataPerPage, err := strconv.Atoi(os.Getenv("DATA_PER_PAGE"))
	if err != nil {
		return news, err
	}

	result, err := repo.elasticConn.ElasticConn.Search().Index(repo.elasticConn.Index).From(dataPerPage * (page - 1)).Size(dataPerPage).Do(context.Background())
	if err != nil {
		return news, err
	}

	hits := result.Hits.Hits
	newsChannel := make(chan newsModel.News, len(hits))
	for _, hit := range hits {
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			return news, err
		}

		go func(newsData []byte, channel chan newsModel.News) {
			newsElastic := newsModel.NewsElasticSearch{}
			err = json.Unmarshal(newsData, &newsElastic)
			if err != nil {
				helper.WriteToLogFile("Failed to Unmarshal News data", err.Error())
			}

			var news newsModel.News
			news, err = repo.GetNewsFromCache(newsElastic)
			if err == redis.Nil {
				news, err = repo.GetNewsDetail(newsElastic)
				if err == nil {
					repo.SaveNewsToCache(news)
				} else if err != nil {
					helper.WriteToLogFile("Failed to get news detail from MySQL", err.Error())
				}
			} else if err != nil {
				helper.WriteToLogFile("Failed to get news detail from cache", err.Error())
			}

			channel <- news
		}(data, newsChannel)
	}

	for i := 0; i < len(hits); i++ {
		value := <-newsChannel
		if value.ID > 0 {
			news = append(news, value)
		}
	}

	sort.SliceStable(news, func(i, j int) bool {
		createdI, err := time.Parse("2006-01-02 15:04:05", news[i].Created)
		if err != nil {
			helper.WriteToLogFile("Failed to parse date time", err.Error())
		}

		createdJ, err := time.Parse("2006-01-02 15:04:05", news[j].Created)
		if err != nil {
			helper.WriteToLogFile("Failed to parse date time", err.Error())
		}

		return createdI.After(createdJ)
	})

	return news, nil
}
