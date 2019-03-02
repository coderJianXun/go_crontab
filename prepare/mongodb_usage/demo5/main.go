package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

// startTime 小于某时间
// $lt: timeStamp
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		delCond    *DeleteCond
		delResult  *mongo.DeleteResult
	)
	// 建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017")); err != nil {
		fmt.Println(err)
		return
	}
	// 选择数据库my_db
	database = client.Database("cron")

	// 选择表my_collection
	collection = database.Collection("log")

	// 删除开始是啊金早于当前时间的所有日志 delete({"timePoint.startTime":{"$lt": 当前时间}})   $lt less than
	delCond = &DeleteCond{beforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	//
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数:", delResult.DeletedCount)
}
