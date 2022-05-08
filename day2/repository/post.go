package repository

import "sync"

type Post struct {
	TopicId int64 `json:"topicId"`
}
type PostDao struct{}

var (
	postIndexMap map[int64][]*Post
	postDao      *PostDao
	postOnce     *sync.Once
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
