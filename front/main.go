package main

import (
	"diplom/front/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//ttpl "github.com/night-codes/go-gin-ttpl"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("config.LoadConfig() err:", err.Error())
		return
	}

	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Static("/css", "./front/templates/css")
	router.Static("/js", "./front/templates/js")
	router.Static("/fronts", "./front/templates/fonts")
	router.Static("/images", "./front/templates/images")
	router.LoadHTMLGlob("front/templates/site/*")
	initializeRoutes()
	//dao.GetAllProducts()
	http.ListenAndServe(config.Config.Port(), router)
	//router.Run()
}
