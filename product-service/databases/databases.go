package databases

import (
	"diplom/product-service/config"
	"fmt"
	"gopkg.in/mgo.v2"
)

// Database shares global database instance
var (
	Database MongoDB
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDbSession *mgo.Session
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	dialUrl := config.Config.MgAddrs()

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDbSession, err = mgo.Dial(dialUrl)

	if err != nil {
		fmt.Println("Can't connect to mongo, go error: ", err)
		return err
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDbSession != nil {
		db.MgDbSession.Close()
	}
}
