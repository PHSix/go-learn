package repository

import (
	"log"
	"sync"
)

type TopicDao struct{}
type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
}

var (
	topicIndexMap map[int64]*Topic
	topicDao      *TopicDao
	topicOnce     sync.Once
	topicMutex    sync.Mutex
	nextId        int64
	idMutex       sync.Mutex
)

// 单例模式，为了节省内存
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}

// 通过topic id 查询topic
func (t TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}

// 创建新的topic
func (t TopicDao) NewTopic(topic *Topic) error {
	topicMutex.Lock()
	topicIndexMap[topic.Id] = topic
	// log
	log.Println("添加了数据,标题为：", topic.Title, topic.Id, topic.Content)
	topicMutex.Unlock()
	return nil
}

func GetNextId() int64 {
	idMutex.Lock()
	nextId++
	idMutex.Unlock()
	return nextId
}
