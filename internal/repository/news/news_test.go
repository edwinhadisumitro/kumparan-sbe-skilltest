package news

import (
	"encoding/json"
	"kumparan-sbe-skilltest/config"
	newsModel "kumparan-sbe-skilltest/internal/model/news"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var newsTestData = newsModel.News{
	ID:      1,
	Author:  "Foo Bar",
	Body:    "Kumparan SBE Skilltest",
	Created: "1970-01-01 00:00:00",
}

var clientRedisMock *redis.Client

func TestGetNewsDetail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "mysql")

	redisClient := new(redis.Client)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id",
		"author",
		"body",
		"created",
	}).AddRow(1, "Foo Bar", "Kumparan SBE Skilltest", time.Now())

	query := "select id, author, body, created from news where id=\\?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	newsRepo := NewNewsRepository(config.NSQPublisher{}, sqlxDB, config.ElasticConn{}, redisClient)
	newsElastic := newsModel.NewsElasticSearch{
		NewsID: 1,
	}
	result, err := newsRepo.GetNewsDetail(newsElastic)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSaveNews(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "mysql")
	redisClient := new(redis.Client)

	defer db.Close()

	query := "insert into news \\(author, body, created\\) values \\(\\?, \\?, \\?\\)"

	mock.ExpectExec(query).WithArgs(newsTestData.Author, newsTestData.Body, newsTestData.Created).WillReturnResult(sqlmock.NewResult(0, 1))

	newsRepo := NewNewsRepository(config.NSQPublisher{}, sqlxDB, config.ElasticConn{}, redisClient)
	_, err = newsRepo.SaveNews(newsTestData)

	assert.NoError(t, err)
}

func NewRedisMock() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	clientRedisMock = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
}

func TestSaveNewsToCache(t *testing.T) {
	NewRedisMock()
	r := redismock.NewNiceMock(clientRedisMock)

	sqlxDB := new(sqlx.DB)

	newsData, _ := json.Marshal(newsTestData)

	r.On("Set", "NEWS:"+strconv.Itoa(newsTestData.ID), string(newsData), 60*time.Second).Return(redis.NewStatusResult("OK", nil))

	newsRepo := NewNewsRepository(config.NSQPublisher{}, sqlxDB, config.ElasticConn{}, clientRedisMock)
	err := newsRepo.SaveNewsToCache(newsTestData)

	assert.NoError(t, err)
}

func TestGetNewsFromCache(t *testing.T) {
	r := redismock.NewNiceMock(clientRedisMock)
	sqlxDB := new(sqlx.DB)

	newsData, _ := json.Marshal(newsTestData)
	newsElasticSearch := newsModel.NewsElasticSearch{
		NewsID: newsTestData.ID,
	}

	r.On("Get", "NEWS:"+strconv.Itoa(newsTestData.ID)).Return(redis.NewStatusResult(string(newsData), nil))

	newsRepo := NewNewsRepository(config.NSQPublisher{}, sqlxDB, config.ElasticConn{}, clientRedisMock)
	news, err := newsRepo.GetNewsFromCache(newsElasticSearch)

	assert.NoError(t, err)
	assert.Equal(t, newsTestData, news)
}
