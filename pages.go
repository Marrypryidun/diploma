package main

import (
	"diplom/dao"
	"diplom/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	obj map[string]interface{}
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
	router.POST("/ajax/product/search", searchProduct)
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

func searchProduct(c *gin.Context) {
	//byteData, err := c.GetRawData()
	//if err != nil {
	//	fmt.Println("pages.go -> searchProduct()->  c.GetRawData() err:", err)
	//	c.JSON(http.StatusBadRequest, obj{"error": err})
	//	return
	//}
	//data := models.ProductSearch{}
	//err = json.Unmarshal(byteData, &data)
	//if err != nil {
	//	fmt.Printf("pages.go -> searchProduct()->  json.Unmarshal() err:%v, data:%s \n", err, string(byteData))
	//	c.JSON(http.StatusBadRequest, obj{"error": err})
	//	return
	//}
	data := models.ProductSearch{}
	if err := c.Bind(&data); err != nil {
		fmt.Printf("pages.go -> searchProduct()-> c.Bind() err:%v, \n", err)
		c.JSON(http.StatusBadRequest, obj{"error": err})
		return
	}
	products := dao.GetProductProductsByName(data)
	c.JSON(http.StatusOK, obj{"products": products})
}
