package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{DB: 10})

	r := gin.Default()
	r.POST("/users/:id/nickname", func(c *gin.Context) {
		key := fmt.Sprintf("users:%s:nickname", c.Param("id"))

		client.Set(key, c.PostForm("nickname"), 0).Result()

		nickname, _ := client.Get(key).Result()
		c.JSON(200, gin.H{
			"nickname": nickname,
		})
	})

	r.GET("users/:id/nickname", func(c *gin.Context) {
		key := fmt.Sprintf("users:%s:nickname", c.Param("id"))
		nickname, _ := client.Get(key).Result()
		c.JSON(200, gin.H{
			"nickname": nickname,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
