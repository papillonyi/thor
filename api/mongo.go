package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/mongo"
	"github.com/papillonyi/thor/mq"
)

func getTrainer(c *gin.Context) {
	name := c.Query("name")
	fmt.Println("try to find: ", name)
	trainer, err := mongo.FindByName(name)
	if err != nil {
		//logging.Error(err)
		c.JSON(500, err)
		return
	}
	fmt.Println("found: ", name)
	c.JSON(200, trainer)
}

func getTest(c *gin.Context) {
	for i := 1; i <= 100; i++ {
		mq.TaskAdd()
	}
	c.JSON(200, "done")
}
