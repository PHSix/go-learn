package repository

import (
	"encoding/json"
	"log"
	"os"
	"strings"
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
	topicIndexMap   map[int64]*Topic
	topicDao        *TopicDao
	topicOnce       sync.Once
	topicMutex      sync.Mutex
	nextId          int64
	idMutex         sync.Mutex
	topicStoreMutex sync.Mutex // 使得不会进行同时的写入操作
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
	go StoreTopic()
	topicMutex.Unlock()
	return nil
}

func GetTopicNextId() int64 {
	idMutex.Lock()
	nextId++
	idMutex.Unlock()
	return nextId
}

func StoreTopic() {
	topicStoreMutex.Lock()
	defer topicStoreMutex.Unlock()
	var items []string
	for _, item := range topicIndexMap {
		c, err := json.Marshal(*item)
		if err != nil {
			log.Println("store topic Marshal failed")
			return
		}
		items = append(items, string(c))
	}
	file, err := os.OpenFile("./data/topic", os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("open topic file failed")
		return
	}
	_, err = file.WriteString(strings.Join(items, "\n"))
	if err != nil {
		log.Println("write topic file failed")
		return
	}
}
