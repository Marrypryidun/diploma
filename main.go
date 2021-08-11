package main

import (
	"diplom/config"
	"diplom/databases"
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
	if err := databases.Database.Init(); err != nil {
		fmt.Println("databases.Database.Init() err:", err.Error())
		return
	}
	defer databases.Database.Close()

	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.Static("/fronts", "./templates/fonts")
	router.Static("/images", "./templates/images")
	router.LoadHTMLGlob("templates/site/*")

	initializeRoutes()
	http.ListenAndServe(":8000", router)
	//router.Run()
}
