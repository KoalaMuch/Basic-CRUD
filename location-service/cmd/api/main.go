package main

import (
	"net/http"
	"basic-crud/data"
	"log"
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort = "80"
	mongoURL = "mongodb://mongo:27017"
)

var client *mongo.Client

type Config struct{
	Models data.Models
}

func main() {

	// connect to mongodb
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	app := Config{
		Models: data.New(client), 
	}

	log.Printf("Starting service on port: \n", webPort)	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connectToMongo() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to MongoDB!\n")
	return c, nil
}