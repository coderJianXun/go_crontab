package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
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

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"` // JobName赋值为job10
}

func main() {
	// mongodb 读取回来的是bson,需要反序列化为LogRecord对象
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		cond       *FindByJobName
		cursor     *mongo.Cursor
		record     *LogRecord
		findopt    options.FindOptions
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

	// 按照jobName字段过滤,想找出jobName= job10 找出5条
	cond = &FindByJobName{JobName: "job10"} // {"jobName: "job10"}

	// 查询 (过滤 + 翻页参数)
	if cursor, err = collection.Find(context.TODO(), cond, findopt.SetSkip(0), findopt.SetLimit(3)); err != nil {
		fmt.Println(err)
		return
	}

	// 反序列化bson到对象
	for cursor.Next(context.TODO()) {
		record = &LogRecord{}
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}

		// 把日志行打印出来
		fmt.Println(*record)
	}

	defer cursor.Close(context.TODO())
}
