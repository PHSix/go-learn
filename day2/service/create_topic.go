package service

import (
	"errors"
	"go-learn/day2/repository"
	"time"
	"unicode/utf16"
)

type createTopicFlow struct {
	Id         int64
	Title      string
	Content    string
	CreateTIme int64
}

func NewCreateTopic(title string, content string) (int64, error) {
	return newCreateTopicFlow(title, content).do()
}

func newCreateTopicFlow(title string, content string) *createTopicFlow {
	return &createTopicFlow{
		Title:   title,
		Content: content,
	}
}
func (t createTopicFlow) checkParam() error {
	if len(utf16.Encode([]rune(t.Content))) < 50 {
		return errors.New("topic content must more than 50")
	}
	return nil
}

func (t *createTopicFlow) do() (int64, error) {

	if err := t.checkParam(); err != nil {
		return 0, err
	}
	topicId := repository.GetTopicNextId()
	t.Id = topicId
	if err := t.create(); err != nil {
		return 0, err
	}
	return t.Id, nil
}

func (t createTopicFlow) create() error {
	topic := &repository.Topic{
		Id:         t.Id,
		Title:      t.Title,
		Content:    t.Content,
		CreateTime: time.Now().Unix(),
	}
	err := repository.NewTopicDaoInstance().NewTopic(topic)
	if err != nil {
		return err
	}

	return nil
}
