package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"go-learn/day2/controller"
	"go-learn/day2/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		// data := controller.QueryPageInfo(topicId)
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/page/add", func(c *gin.Context) {
		bodyReader := bufio.NewReader(c.Request.Body)
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			c.JSON(400, gin.H{})
			return
		}
		data := controller.CreateTopic(body)

		c.JSON(200, data)
	})
	err := r.Run(":5000")
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		log.Println("init repository failed")
		return err
	}
	return nil
}
