package main

import "github.com/gin-gonic/gin"

var router = gin.Default()

// router := gin.Default()

func init() {
	router.GET("/content/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "get请求"})
	})

	router.POST("/content", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "post请求"})
	})

	router.DELETE("/content/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "删除id"})
	})

	router.PUT("/content/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "更新id"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
