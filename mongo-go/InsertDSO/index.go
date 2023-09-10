package main

import (
	"context"
	"log"
	"time"

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

		// Insert
		/*
			volvo := Vehicle{Name: "Volvo Extreme", Speed: 210}
			if result, err := collection.InsertOne(ctx, volvo); err != nil {
				panic("Something went wrong: " + err.Error())
			} else {
				log.Println("Single record has been saved!")
				log.Println(result)
			}
		*/

		vehicles := []interface{}{
			Vehicle{Name: "Toyota", Speed: 200},
			Vehicle{Name: "Toyota Extreme", Speed: 220},
		}
		if result, err := collection.InsertMany(ctx, vehicles); err != nil {
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
