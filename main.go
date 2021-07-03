package main

import "github.com/gin-gonic/gin"



func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.LoadHTMLFiles("./templates/index.html")
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.Static("/fronts", "./templates/fronts")
	router.Static("/images", "./templates/images")
	initializeRoutes()
	router.Run()
}