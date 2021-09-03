package dao

import (
	"diplom/product-service/config"
	"diplom/product-service/databases"
	"diplom/product-service/models"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type ProductDaoModuleStruct struct {
	collectionName string
}

var ProductDaoModule = ProductDaoModuleStruct{
	collectionName: "products",
}

func (m *ProductDaoModuleStruct) GetAllProducts() []models.Product {

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(config.Config.MgDbName()).C(m.collectionName)

	var products []models.Product
	err := collection.Find(bson.M{}).All(&products)
	if err != nil {
		fmt.Println("product.go -> GetAllProducts() -> collection.Find err:", err)
	}
	return products
}

func (m *ProductDaoModuleStruct) GetProductProductsByName(data models.ProductSearch) []models.Product {

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(config.Config.MgDbName()).C("products")

	var products []models.Product
	search := bson.D{}
	if len(data.Product) > 0 {
		search = bson.D{bson.DocElem{"name", data.Product}}
	}
	err := collection.Find(search).Skip(0).Limit(3).All(&products)
	if err != nil {
		fmt.Printf("product.go->GetProductProductsByName()->Find() err:%v, name:%s\n", err, data.Product)
		return []models.Product{}
	}
	return products
}
