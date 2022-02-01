package controllers

import (
	"diplom/product-service/dao"
	"diplom/product-service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
)

type ProductModuleStruct struct {
	ProductCache map[string]models.Product
	sync.RWMutex
}

// router
var (
	Router        *gin.Engine
	ProductModule = ProductModuleStruct{
		ProductCache: make(map[string]models.Product),
	}
)

func (m *ProductModuleStruct) LoadCache() {
	products := dao.ProductDaoModule.GetAllProducts()
	ProductModule.Lock()
	for k := range products {
		ProductModule.ProductCache[products[k].Id] = products[k]
	}
	ProductModule.Unlock()
}

func (m *ProductModuleStruct) InitializeRoutes() {
	Router.POST("/getProductsByName", m.getProductsByNameHandler)
}

func (m *ProductModuleStruct) getProductsByNameHandler(c *gin.Context) {
	inputData := models.ProductSearch{}

	if err := c.Bind(&inputData); err != nil {
		fmt.Println("product.go ->getProductsByNameHandler() -> c.Bind() err = ", err)
		c.Status(http.StatusBadRequest)
		return
	}

	products := m.getProductByNameFromCache(inputData)
	c.JSON(http.StatusOK, products)
}

func (m *ProductModuleStruct) getProductByNameFromCache(search models.ProductSearch) []models.Product {
	products := []models.Product{}
	m.RLock()
	for _, v := range m.ProductCache {
		if strings.Contains(v.Name, search.Product) {
			products = append(products, v)
		}
		if len(products) == search.Number {
			break
		}
	}
	m.RUnlock()
	return products
}
