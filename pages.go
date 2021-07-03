package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router
var router *gin.Engine

func initializeRoutes() {
	router.GET("/", showIndexPage)
}

func showIndexPage(c * gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}