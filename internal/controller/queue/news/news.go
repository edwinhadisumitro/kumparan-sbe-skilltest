package news

import (
	"encoding/json"
	"kumparan-sbe-skilltest/helper"
	newsInterface "kumparan-sbe-skilltest/internal/interface/news"
	newsModel "kumparan-sbe-skilltest/internal/model/news"

	"github.com/nsqio/go-nsq"
)

// NSQHandler : Struct for implementing NSQ Handler interface. No controller interface needed.
type NSQHandler struct {
	newsLibrary newsInterface.Library
}

// NewNSQHandler : Establish new NSQ Handler
func NewNSQHandler(newsLibrary newsInterface.Library) nsq.Handler {
	return &NSQHandler{
		newsLibrary: newsLibrary,
	}
}

// HandleMessage : Handle Incoming NSQ Message
func (h *NSQHandler) HandleMessage(m *nsq.Message) error {
	var err error

	if len(m.Body) > 0 {
		var news newsModel.News
		err = json.Unmarshal(m.Body, &news)
		if err != nil {
			return err
		}

		err = h.newsLibrary.SaveNews(news)
		if err != nil {
			helper.WriteToLogFile("Failed to save news", err.Error())
			return err
		}
	}

	return nil
}
