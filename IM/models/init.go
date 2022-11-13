package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
var Mongo=InitMongo()
func InitMongo() *mongo.Database{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://42.193.120.85:27017"))
	if err!=nil{
		log.Println("Connect Mongo Failed",err)
	}
	return client.Database("im")
}
