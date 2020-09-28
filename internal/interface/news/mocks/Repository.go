// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	news "kumparan-sbe-skilltest/internal/model/news"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetNews provides a mock function with given fields: page
func (_m *Repository) GetNews(page int) ([]news.News, error) {
	ret := _m.Called(page)

	var r0 []news.News
	if rf, ok := ret.Get(0).(func(int) []news.News); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]news.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewsDetail provides a mock function with given fields: newsElastic
func (_m *Repository) GetNewsDetail(newsElastic news.NewsElasticSearch) (news.News, error) {
	ret := _m.Called(newsElastic)

	var r0 news.News
	if rf, ok := ret.Get(0).(func(news.NewsElasticSearch) news.News); ok {
		r0 = rf(newsElastic)
	} else {
		r0 = ret.Get(0).(news.News)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.NewsElasticSearch) error); ok {
		r1 = rf(newsElastic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNewsFromCache provides a mock function with given fields: newsElastic
func (_m *Repository) GetNewsFromCache(newsElastic news.NewsElasticSearch) (news.News, error) {
	ret := _m.Called(newsElastic)

	var r0 news.News
	if rf, ok := ret.Get(0).(func(news.NewsElasticSearch) news.News); ok {
		r0 = rf(newsElastic)
	} else {
		r0 = ret.Get(0).(news.News)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.NewsElasticSearch) error); ok {
		r1 = rf(newsElastic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublishNews provides a mock function with given fields: _a0
func (_m *Repository) PublishNews(_a0 news.News) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(news.News) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveNews provides a mock function with given fields: _a0
func (_m *Repository) SaveNews(_a0 news.News) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(news.News) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.News) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveNewsID provides a mock function with given fields: _a0
func (_m *Repository) SaveNewsID(_a0 news.NewsElasticSearch) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(news.NewsElasticSearch) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveNewsToCache provides a mock function with given fields: _a0
func (_m *Repository) SaveNewsToCache(_a0 news.News) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(news.News) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}