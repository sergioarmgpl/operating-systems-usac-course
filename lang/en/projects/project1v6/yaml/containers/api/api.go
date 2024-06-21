package main

import (
    "net/http"
    "context"

    "github.com/redis/go-redis/v9"
    "github.com/gin-gonic/gin"
)

type message struct {
    Message     string  `json:"message"`
    Country string  `json:"country"`
}

type result struct {
    Processed     string  `json:"processed"`
}

func main() {
    router := gin.Default()
	router.GET("/_health", getHealth)
    router.POST("/message", postMessage)
    router.Run(":1234")
}

func getHealth(c *gin.Context) {
	var   newResult result
	newResult.Processed = "done"
    c.IndentedJSON(http.StatusOK,newResult)
}

func postMessage(c *gin.Context) {
    var newMessage message
    if err := c.BindJSON(&newMessage); err != nil {
        return
    }

    var ctx = context.Background()

    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    err := rdb.HIncrBy(ctx,"countries",newMessage.Country,1).Err()
    if err != nil {
        panic(err)
    }

    c.IndentedJSON(http.StatusCreated, newMessage)
}
