package news

import (
	"encoding/json"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
)

func (repo *newsRepository) PublishNews(news newsModel.News) error {
	payload, err := json.Marshal(news)
	if err != nil {
		return err
	}

	err = repo.nsqProducer.Producer.Publish(repo.nsqProducer.Topic, payload)
	if err != nil {
		return err
	}

	return nil
}
