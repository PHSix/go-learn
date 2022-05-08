package repository

import "sync"

type TopicDao struct{}
type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

var (
	topicIndexMap map[int64]*Topic
	topicDao      *TopicDao
	topicOnce     *sync.Once
)

// 单例模式，为了节省内存
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}

func (t TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
