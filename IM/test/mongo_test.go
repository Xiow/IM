package test

import (
	"IM/models"
	"context"
	_ "context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestFindone(t *testing.T)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://42.193.120.85:27017"))
	if err!=nil{
		t.Fatal(err)
	}
	db:=client.Database("im")
	ub:=new(models.UserBasic)
	err=db.Collection("user_basic").FindOne(context.Background(),bson.D{}).Decode(ub)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println("ub====>",ub)

}
func TestFind(t *testing.T)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://42.193.120.85:27017"))
	if err!=nil{
		t.Fatal(err)
	}
	db:=client.Database("im")
	cursor, err := db.Collection("user_room").Find(context.Background(), bson.D{})
	urs := make([]*models.UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(models.UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}
	for _, v := range urs {
		fmt.Println("UserRoom ==> ", v)
	}
}

