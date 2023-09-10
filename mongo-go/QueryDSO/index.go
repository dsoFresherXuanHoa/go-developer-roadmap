package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Vehicle struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Speed int                `bson:"speed"`
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

		/*
			var vehicle Vehicle
			var filter = bson.D{}
			if err := collection.FindOne(ctx, filter).Decode(&vehicle); err != nil {
				panic("Something went wrong: " + err.Error())
				} else {
					log.Println(reflect.TypeOf(vehicle))
					if data, err := json.Marshal(vehicle); err != nil {
						panic("Can't to parse to json: " + err.Error())
						} else {
							log.Println(string(data))
						}
					}
		*/
		var filter = bson.D{}
		if cur, err := collection.Find(ctx, filter); err != nil {
			panic("Something went wrong: " + err.Error())
		} else {
			var vehicles []Vehicle
			if err := cur.All(ctx, &vehicles); err != nil {
				panic("Can't parse: " + err.Error())
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
