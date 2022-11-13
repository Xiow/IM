package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Identity  string `bson:"identity"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"create_at"`
	UpdatedAt int64  `bson:"update_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}


func GetUserBasicBYAccountPassword(account string,password string)(*UserBasic,error){

	ub:=new(UserBasic)
	err:=Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(),bson.D{{"account",account},{"password",password}}).Decode(ub)
	return ub,err
}
//获取用户详细信息
func GetUserBasicByIdentity(identity string)(*UserBasic,error){
	ub:=new(UserBasic)
	err:=Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(),bson.D{{"identity",identity}}).Decode(ub)
	return ub,err
}
//根据邮箱获取用户个数
func GetUserBasicCountByEmail(email string)(int64,error){
	return Mongo.Collection(UserBasic{}.CollectionName()).CountDocuments(context.Background(),bson.D{{"email",email}})
}
