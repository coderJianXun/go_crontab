package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell 命令
	Err       string    `bson:"err"`       // J脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		result     *mongo.InsertOneResult
		docId      primitive.ObjectID
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

	// 插入数据
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix()},
	}

	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}

	// id: 默认生成一个全局唯一的ID,ObjectID:12字节的二进制
	docId = result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
}
