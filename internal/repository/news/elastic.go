package news

import (
	"context"
)

func (repo *newsRepository) SaveNewsID(id int) error {
	var err error
	data := make(map[string]int)
	data["news_id"] = id

	_, err = repo.elasticConn.ElasticConn.Index().Index(repo.elasticConn.Index).BodyJson(data).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
