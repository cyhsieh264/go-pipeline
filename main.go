package main

import (
	"github.com/cyhsieh264/go-pipeline/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	r.GET("/foo", func(c *gin.Context) {
		client := service.New()

		c.JSON(200, gin.H{
			"foo": client.Foo(),
		})
	})

	r.Run(":8080")
}
