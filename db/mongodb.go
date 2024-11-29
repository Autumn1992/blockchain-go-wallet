package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"walletserver/log"
)

// 游戏金币详情

var gMongodbClient *mongo.Client

func InitMongodb() {
	return
	//fmt.Println(GetConfig().Mongod.Host)
	moptions := options.Client().ApplyURI(GetConfig().Mongod.Host)
	moptions.SetMaxPoolSize(10)
	moptions.SetMinPoolSize(5)
	moptions.SetMaxConnIdleTime(10 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, moptions)
	if err != nil {
		panic(err)
	}
	var readf readpref.ReadPref
	err = client.Ping(context.Background(), &readf)
	if err != nil {
		log.Error("无法连接到mongodb!")
		panic(err)
	}

	gMongodbClient = client

}

type Person struct {
	Name string
	Age  int
}

// test
func AddDataTest() {
	// 选择数据库和集合
	collection := gMongodbClient.Database("mydatabase").Collection("people")
	ctx := context.TODO()
	// 插入数据
	_, err := collection.InsertOne(ctx, Person{"Alice", 30})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("插入成功")
}
func GetMongodbClient() *mongo.Client {
	return gMongodbClient
}
