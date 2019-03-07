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
		logArr     []interface{}
		collection *mongo.Collection
		record     *LogRecord
		insertId   interface{}
		result     *mongo.InsertManyResult
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

	// 插入记录(bson)
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	// 批量插入多条
	logArr = []interface{}{record, record, record}

	// 发起插入
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}

	for _, insertId = range result.InsertedIDs {
		docId = insertId.(primitive.ObjectID)
		fmt.Println("自增ID:", docId.Hex())
	}
}
