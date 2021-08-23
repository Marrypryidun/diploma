package dao

import (
	"diplom/config"
	"diplom/databases"
	"diplom/models"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func GetAllProducts() ([]models.Product, error) {

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(config.Config.MgDbName()).C("products")

	var products []models.Product
	err := collection.Find(bson.M{}).All(&products)
	return products, err
}

func GetProductProductsByName(data models.ProductSearch) []models.Product {

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(config.Config.MgDbName()).C("products")

	var products []models.Product
	search := bson.D{}
	if len(data.Product) > 0 {
		search = bson.D{bson.DocElem{"name",  data.Product}}
	}
	err := collection.Find(search).Skip(0).Limit(3).All(&products)
	if err != nil {
		fmt.Printf("product.go->GetProductProductsByName()->Find() err:%v, name:%s\n", err, data.Product)
		return []models.Product{}
	}
	return products
}
