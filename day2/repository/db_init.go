package repository

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// 初始化topic
func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(open)
	topicTmp := make(map[int64]*Topic)
	for scan.Scan() {
		text := scan.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmp[topic.Id] = &topic
		if topic.Id > nextId {
			nextId = topic.Id
		}
	}
	topicIndexMap = topicTmp
	return nil
}

// 初始化post
func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(open)
	var postTmp = make(map[int64][]*Post)
	for scan.Scan() {
		text := scan.Text()
		var post Post
		if err = json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postTmp[post.TopicId] = append(postTmp[post.TopicId], &post)
	}
	postIndexMap = postTmp
	return nil
}

func Init(filePath string) error {
	if err := initPostIndexMap(filePath); err != nil {
		log.Println("init post map failed")
		return err
	}
	if err := initTopicIndexMap(filePath); err != nil {
		log.Println("init topic map failed")
		return err
	}
	return nil
}
