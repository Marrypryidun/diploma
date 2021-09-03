package controlers

import (
	"diplom/front/models"
	"diplom/front/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	obj map[string]interface{}

	ProductModuleStruct struct {
		serviceUrl string
	}
)

var ProductModule = ProductModuleStruct{
	serviceUrl: "http://localhost:8010",
}

func (m *ProductModuleStruct) SearchProduct(c *gin.Context) {
	data := models.ProductSearch{}
	if err := c.Bind(&data); err != nil {
		fmt.Printf("pages.go -> searchProduct()-> c.Bind() err:%v, \n", err)
		c.JSON(http.StatusBadRequest, obj{"error": err})
		return
	}
	data.Number = 4
	answerByte, ok := utils.RequestToAPI(http.MethodPost, m.serviceUrl+"/getProductsByName", data)
	if !ok {
		fmt.Println("product.go -> SearchProduct() -> utils.RequestToAPI() err request to product service.")
		c.JSON(http.StatusBadRequest, obj{"error": "!ok"})
		return
	}
	fmt.Printf("product.go -> SearchProduct() answerByte:%v \n", string(answerByte))
	products := []models.Product{}
	err := json.Unmarshal(answerByte, &products)
	if err != nil {
		fmt.Println("product.go -> SearchProduct() -> json.Unmarshal() err:", err)
		c.JSON(http.StatusBadRequest, obj{"error": err})
		return
	}
	c.JSON(http.StatusOK, obj{"products": products})
}
