package db

import (
	`fmt`

	`github.com/sushshring/K8_Sample/common`
	`github.com/globalsign/mgo/bson`
	`github.com/go-bongo/bongo`
)

type DB struct {
	config *bongo.Config
	connection *bongo.Connection
}

func (db DB) Connect() DB{
	db.config = &bongo.Config{
		ConnectionString: common.GetEnv("MONGO_DB_HOSTNAME", "localhost"),
		Database: common.GetEnv("MONGO_DB_NAME", "bongotest"),
	}
	connection, err := bongo.Connect(db.config)
	if err != nil {
		fmt.Println(err)
	}
	db.connection = connection
	return db
}

func (db DB) AddObject(data bongo.Document, collectionName string) error {
	return db.connection.Collection(collectionName).Save(data)
}

func (db DB) GetObject(id bson.ObjectId, collectionName string, objectToReturn *bongo.Document) error {
	return db.connection.Collection(collectionName).FindById(id, *objectToReturn)
}

func (db DB) GetAllObjects(collectionName string) *bongo.ResultSet{
	return db.connection.Collection(collectionName).Find(nil)
}