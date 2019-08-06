package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client *mongo.Client
	DB     *mongo.Database
	err    error
)

func Setup() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	opt := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	opt.SetMaxPoolSize(20)
	client, err = mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("runoob")

	fmt.Println("Connected to MongoDB!")

}
