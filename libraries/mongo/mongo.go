package mongo

import (
	"context"
	"fmt"
	"log"
	"time"
	"work_report/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgoClient *mongo.Client

type MgoCondition struct {
	Filter     bson.M
	PagePerNum int64
	PageNum    int64
}

func init() {
	connect()
}

func connect() {
	mgoConf := config.GMConfig["system_log"]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mgoConf.User, mgoConf.Pass, mgoConf.Host, mgoConf.Port)
	var err error
	MgoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("connnect mongo error :" + err.Error())
	}
}

func InsertOne(db string, table string, data bson.M) error {
	_, err := MgoClient.Database(db).Collection(table).InsertOne(context.Background(), data)
	return err
}

func FindOne(db string, table string, filter bson.M) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := MgoClient.Database(db).Collection(table).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func Find(db string, table string, condition MgoCondition) (interface{}, error) {
	var results []map[string]interface{}
	var skipNum int64
	opt := options.Find()
	fmt.Println(condition)
	if condition.PagePerNum != 0 {
		skipNum = 0
		if condition.PageNum > 1 {
			skipNum = (condition.PageNum - 1) * condition.PagePerNum
		}
		opt.SetLimit(condition.PagePerNum)
		if skipNum > 0 {
			opt.SetSkip(skipNum)
		}
		fmt.Println(skipNum)
	}
	cur, err := MgoClient.Database(db).Collection(table).Find(context.TODO(), condition.Filter, opt)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem map[string]interface{}
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}

func UpdateOne(db string, table string, filter bson.D, updateData bson.D) (int64, error) {
	updateMatch, err := MgoClient.Database(db).Collection(table).UpdateOne(context.TODO(), filter, updateData)
	return updateMatch.UpsertedCount, err
}
