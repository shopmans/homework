package pkg

import (
	// 标准包
	"context"
	"fmt"
	"os"
	"strconv"

	// 第三方包
	"github.com/jinzhu/configor"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// 内部包
)

type DbConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"root"`
	Password     string `default:"root"`
	DatabaseName string `default:"rbiot"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

var db *mongo.Database

func NewDB() *mongo.Database {
	if nil != db {
		return nil
	}

	cf, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	mongoConfig := cf + "/configs/mongo.yml"

	dbConfig := &DbConfig{}
	configor.Load(dbConfig, mongoConfig)

	param := "mongodb://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(param))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	db = client.Database(dbConfig.DatabaseName)

	return db
}

func GetDB() *mongo.Database {
	return db
}
