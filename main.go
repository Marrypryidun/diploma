package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//ttpl "github.com/night-codes/go-gin-ttpl"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.Static("/fronts", "./templates/fonts")
	router.Static("/images", "./templates/images")
	router.LoadHTMLGlob("templates/site/*")

	initializeRoutes()
	http.ListenAndServe(":8001", router)
	//router.Run()
}
