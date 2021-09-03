package main

import (
	"diplom/product-service/config"
	"diplom/product-service/controllers"
	"diplom/product-service/databases"
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

	controllers.Router = gin.New()
	controllers.ProductModule.InitializeRoutes()
	controllers.ProductModule.LoadCache()
	//dao.GetAllProducts()
	http.ListenAndServe(config.Config.Port(), controllers.Router)
	//router.Run()
}
