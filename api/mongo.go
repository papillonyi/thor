package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/mongo"
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
