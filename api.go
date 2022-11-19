package main

import (
	"cmsGo/app/content"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// router := gin.Default()

func init() {
	router.GET("/content/:id", content.GET)

	router.POST("/content", content.POST)

	router.DELETE("/content/:id", content.DELETE)

	router.PUT("/content/:id", content.PUT)
}
