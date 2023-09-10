package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		panic("Can't connect to mongodb services: " + err.Error())
	} else {
		log.Println("Connection has been create: ")
		// Checking connection:
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			panic("Don't has any connection: " + err.Error())
		} else {
			log.Println("Connection has been found!")
		}

		// Insert
		collection := client.Database("mongo_go_dev").Collection("vehicles")
		if res, err := collection.InsertOne(ctx, bson.D{{"name", "Volvo"}, {"speed", 200}}); err != nil {
			panic("Can't insert record into database!")
		} else {
			log.Println("Record has been saved: ", res.InsertedID)
		}

		// Query
		if cur, err := collection.Find(ctx, bson.D{}); err != nil {
			panic("Can't query record from database!")
		} else {
			defer cur.Close(ctx)
			for cur.Next(ctx) {
				var result bson.D
				if err := cur.Decode(&result); err != nil {
					panic("Something went wrong while decoding record!")
				} else {
					log.Println(result)
				}
			}
		}

		// Disconcertion:
		defer func() {
			if err := client.Disconnect(ctx); err != nil {
				panic("Something went wrong: " + err.Error())
			} else {
				log.Println("Connection has been disconnected!")
			}
		}()
	}
}
