package mq

import (
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func TaskAdd() *result.AsyncResult {
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
	asyncResult, err := server.SendTask(signature)
	if err != nil {
	}

	return asyncResult
}
