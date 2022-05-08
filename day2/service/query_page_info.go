package service

import (
	"errors"
	"go-learn/day2/repository"
	"sync"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type queryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *repository.Topic
	posts    []*repository.Post
}

func newQueryPageInfoFlow(id int64) *queryPageInfoFlow {
	return &queryPageInfoFlow{
		topicId: id,
	}
}

func (f queryPageInfoFlow) checkParam() error {
	if f.topicId < 0 {
		return errors.New("topic id must bigger than 0")
	}
	return nil
}

func (f *queryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostByTopicId(f.topicId)
		f.posts = posts
	}()

	wg.Wait()
	return nil

}

func (f *queryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}

func (f *queryPageInfoFlow) do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return newQueryPageInfoFlow(topicId).do()
}
