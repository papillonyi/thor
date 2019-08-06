package mongo

import (
	"context"
	"fmt"
	"github.com/papillonyi/thor/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func AddTrainer(name string, age int, city string) Trainer {
	trainer := Trainer{name, age, city}
	collection := DB.Collection("trainers")
	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		logging.Error(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return trainer
}

func FindByName(name string) (trainer Trainer, err error) {
	filter := bson.D{{"name", name}}
	err = DB.Collection("trainers").FindOne(context.Background(), filter).Decode(&trainer)
	//if err != nil {
	//	logging.Error("err")
	//}
	return
}
