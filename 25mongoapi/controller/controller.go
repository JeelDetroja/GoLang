package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/jeeldetroja/GoLang/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://jeeldetroja:snOalVRSvNGlKRm1@cluster0.a0lsvse.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB conection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instanc is ready")
}

//MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert 1 movie in db with id: ", inserted.InsertedID)
}

//update 1 record
