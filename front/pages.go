package main

import (
	"diplom/front/controlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// router
var router *gin.Engine

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/about-us.html", showAboutUsPage)
	router.GET("/index.html", showIndexPage)
	router.GET("/news.html", showNewsPage)
	router.GET("/contact-us.html", showContactUsPage)
	router.GET("body-visualizer.html", showBodyVisualizer)
	router.GET("product.html", showProductPage)
	router.POST("/ajax/product/search", controlers.ProductModule.SearchProduct)
}

func showIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func showAboutUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.html", nil)
}

func showNewsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "news.html", nil)
}
func showContactUsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "contact-us.html", nil)
}

func showBodyVisualizer(c *gin.Context) {
	c.HTML(http.StatusOK, "body-visualizer.html", nil)
}

func showProductPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product.html", nil)
}
