package controller

import (
	"encoding/json"
	"go-learn/day2/service"
)

type TopicResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type TopicRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateTopic(body []byte) *TopicResponse {
	var requestBody TopicRequestBody
	err := json.Unmarshal(body, &requestBody)
	if err != nil {
		return &TopicResponse{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	_, err = service.NewCreateTopic(requestBody.Title, requestBody.Content)

	if err != nil {
		return &TopicResponse{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &TopicResponse{
		Code: 0,
		Msg:  "success",
	}
}
