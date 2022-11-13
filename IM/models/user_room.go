package models

import (
	//"Ginchat/models"
	"context"
	//"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type UserRoom struct {

	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	//RoomType     int    `bson:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt    int64  `bson:"created_at"`
	UpdatedAt    int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}
//func GetUserRoomByUserIdentityRoomIdntity(userIdentity,roomIdentity string)(*UserRoom,error){
//
//	ur:=new(UserRoom)
//	err:=Mongo.Collection(UserRoom{}.CollectionName()).FindOne(context.Background(),
//		bson.D{{"user_identity",userIdentity},{"room_identity",roomIdentity}}).Decode(ur)
//	return ur, err
//
//}
func GetUserRoomByUserIdentityRoomIdntity(userIdentity string)(*UserRoom,error){
	ur:=new(UserRoom)
	err:=Mongo.Collection(UserRoom{}.CollectionName()).FindOne(context.Background(),
		bson.D{{"user_identity",userIdentity}}).Decode(ur)
	return ur, err
}
func GetUserRoomByRoomidentity(roomIdentity string)([]*UserRoom,error) {

	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"user_identity",roomIdentity}})
	if err!=nil{
		return nil,err
	}
	urs := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Fatal(err)
		}
		urs = append(urs, ur)
	}
	//for _, v := range urs {
	//	fmt.Println("UserRoom ==> ", v)
	//}
	return urs,nil
}
