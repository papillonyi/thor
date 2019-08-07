package mq

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
	"time"
)

var cnf = &config.Config{
	Broker:        "amqp://guest:guest@localhost:5672/",
	DefaultQueue:  "machinery_tasks",
	ResultBackend: "mongodb://root:root@localhost:27017/",
	MongoDB: &config.MongoDBConfig{
		Database: "taskresult",
	},
	AMQP: &config.AMQPConfig{
		Exchange:     "machinery_exchange",
		ExchangeType: "direct",
		BindingKey:   "machinery_task",
	},
}

var (
	err    error
	server *machinery.Server
)

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	time.Sleep(10 * time.Second)
	return sum, nil
}

func Multiply(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	return sum, nil
}

func Setup() {
	server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("add server")

	//err = server.RegisterTasks(map[string]interface{}{
	//	"add":      Add,
	//	"multiply": Multiply,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	log.Printf("add task")

	for i := 1; i <= 100000; i++ {
		signature := &tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 1,
				},
			},
		}

		//fmt.Println(i)
		_, err := server.SendTask(signature)

		//taskState := asyncResult.GetState()
		//fmt.Printf("Current state of %v task is:\n", taskState.TaskUUID)
		//fmt.Println(taskState.State)

		if err != nil {
		}
	}

}

func SetupWork() {
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

	worker := server.NewWorker("worker_name", 5)
	err = worker.Launch()
	if err != nil {
		log.Fatal(err)
	}
}
