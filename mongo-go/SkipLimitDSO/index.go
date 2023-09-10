package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		collection := client.Database("mongo_go_dev").Collection("vehicles")

		// Skip, Limit
		var filter = bson.D{{}}
		var options = options.Find().SetSkip(1).SetLimit(2)
		if cur, err := collection.Find(ctx, filter, options); err != nil {
			panic("Something went wrong: " + err.Error())
		} else {
			var vehicles []Vehicle
			if err := cur.All(ctx, &vehicles); err != nil {
				panic("Something went wrong: " + err.Error())
			} else {
				for _, v := range vehicles {
					if data, err := json.Marshal(v); err != nil {
						panic("Can't parse: " + err.Error())
					} else {
						log.Println(string(data))
					}
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
