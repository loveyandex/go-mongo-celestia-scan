package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DBName = "tia-db"

//GetMongoDbConnection get connection of mongodb
var Client, err = GetMongoDbConnection()

func GetMongoDbConnection() (*mongo.Client, error) {

	dbserver := os.Getenv("MONGODBHOST")
	URIII := "mongodb://" + dbserver + "/?replicaSet=myReplicaSet"
	if dbserver == "" {
		dbserver = "localhost"
		URIII = "mongodb://" + dbserver + ":27017/"

	}
	fmt.Println("dbserver ", dbserver)
	fmt.Println("uriii ", URIII)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URIII))
	// fmt.Printf("db client after connection %v ", client)

	if err != nil {
		fmt.Printf("err %v \n", err)
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Printf("err client.Ping %v \n", err)
		log.Fatal(err)
	}
	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	if err != nil {
		return nil, err
	}

	collection := Client.Database(DbName).Collection(CollectionName)

	return collection, nil
}

func Collection(CollectionName string) *mongo.Collection {
	// fmt.Printf("robinClient %v \n", robinClient)
	collection := Client.Database(DBName).Collection(CollectionName)
	return collection
}
