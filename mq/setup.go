package mq

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/papillonyi/thor/pkg/setting"
	"log"
)

var (
	err    error
	server *machinery.Server
)

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func Multiply(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	return sum, nil
}

func getMqCnf() *config.Config {
	return &config.Config{
		Broker: fmt.Sprintf("amqp://%s:%s@%s/",
			setting.AmqpSetting.User,
			setting.AmqpSetting.Password,
			setting.AmqpSetting.Host,
		),
		DefaultQueue: "machinery_tasks",
		ResultBackend: fmt.Sprintf("mongodb://%s:%s@%s/",
			setting.MongoSetting.User,
			setting.MongoSetting.Password,
			setting.MongoSetting.Host,
		),
		MongoDB: &config.MongoDBConfig{
			Database: "taskresult",
		},
		AMQP: &config.AMQPConfig{
			Exchange:     "machinery_exchange",
			ExchangeType: "direct",
			BindingKey:   "machinery_task",
		},
	}
}

func Setup() {
	var cnf = getMqCnf()

	server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}

	//for i := 1; i <= 100000; i++ {
	//	signature := &tasks.Signature{
	//		Name: "add",
	//		Args: []tasks.Arg{
	//			{
	//				Type:  "int64",
	//				Value: 1,
	//			},
	//			{
	//				Type:  "int64",
	//				Value: 1,
	//			},
	//		},
	//	}
	//	_, err := server.SendTask(signature)
	//	if err != nil {
	//	}
	//}

}

func SetupWork() {
	var cnf = getMqCnf()

	server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("add server")

	err = server.RegisterTasks(map[string]interface{}{
		"add":      Add,
		"multiply": Multiply,
	})
	if err != nil {
		log.Fatal(err)
	}

	worker := server.NewWorker("worker_name", 100)
	err = worker.Launch()
	if err != nil {
		log.Fatal(err)
	}
}
