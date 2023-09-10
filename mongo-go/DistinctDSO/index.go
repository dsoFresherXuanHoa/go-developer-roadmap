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

type Vehicle struct {
	Name  string
	Speed int
}

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

		// Distinct:
		collection := client.Database("mongo_go_dev").Collection("vehicles")
		var filter = bson.D{{"name", "Toyota Extreme"}}
		if result, err := collection.Distinct(ctx, "name", filter); err != nil {
			panic("Something went wrong: " + err.Error())
		} else {
			log.Println(result)
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
