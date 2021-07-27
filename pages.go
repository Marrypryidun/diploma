package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router
var router *gin.Engine

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/about-us.html", showAboutUsPage)
	router.GET("/index.html", showIndexPage)
	router.GET("/typography.html", showTypographyPage)
	router.GET("/contact-us.html", showContactUsPage)
	router.GET("body-visualizer.html", showBodyVisualizer)
}

func showIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func showAboutUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.html", nil)
}

func showTypographyPage(c *gin.Context) {
	c.HTML(http.StatusOK, "typography.html", nil)
}
func showContactUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact-us.html", nil)
}

func showBodyVisualizer(c *gin.Context) {
	c.HTML(http.StatusOK, "body-visualizer.html", nil)
}
