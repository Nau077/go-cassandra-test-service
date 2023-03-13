package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	model "github.com/Nau077/cassandra-golang-sv/internal/presentation/model"
	"github.com/gin-gonic/gin"
)

type HTTPPostSender interface {
	InsertData(string, string)
	GetByID(int64)
}

type HTTPPost struct{}

func getId() chan interface{} {
	c := make(chan interface{})
	requestURL := "https://jsonplaceholder.typicode.com/posts"
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer close(c)

		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			fmt.Printf("client: could not create request: %s\n", err)
			c <- "coud not get"
		}
		res, err := http.DefaultClient.Do(req)
		var data model.Posts

		resBody, err := ioutil.ReadAll(res.Body)
		if err := json.Unmarshal(resBody, &data); err != nil {
			fmt.Println("failed to unmarshal:", err)
		} else {
			fmt.Println(data[0]) // 5
		}
		c <- data

	}()

	return c

}

func (h *HTTPPost) GetByID(c *gin.Context) {

	fmt.Println(c.Query("id"))

	select {
	case posts := <-getId():
		c.JSON(200, gin.H{
			"users": posts,
		})
	}
}

func (h *HTTPPost) InsertData(c *gin.Context) {
	return
}

func NewPostHTTPHandler(r *gin.Engine) {
	handler := &HTTPPost{}

	api := r.Group("/post")
	api.GET("/", handler.GetByID)
	api.POST("/add", handler.InsertData)

}
