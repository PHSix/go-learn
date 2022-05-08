package repository

import (
	"sync"
)

type Post struct {
	TopicId    int64  `json:"topicId"`
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
}
type PostDao struct{}

var (
	postIndexMap map[int64][]*Post
	postDao      *PostDao
	postOnce     sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(func() {
		postDao = &PostDao{}
	})
	return postDao
}

func (p PostDao) QueryPostByTopicId(id int64) []*Post {
	return postIndexMap[id]
}
