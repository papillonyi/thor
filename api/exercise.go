package api

import (
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/exercise"
	"log"
)

func addResource(manager exercise.StoreManager, c *gin.Context) {
	err := c.ShouldBindJSON(manager)
	if err != nil {
		log.Print(err)
		c.JSON(500, err)
		return
	}

	err = manager.Add()
	if err != nil {
		log.Print(err)
		c.JSON(500, err)
		return
	}
	c.JSON(200, manager)
	return
}

func getResource(manager exercise.StoreManager, c *gin.Context) {
	namespace := c.Query("namespace")
	actionList, err := manager.GetListByNamespace(namespace)
	if err != nil {
		log.Print(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, actionList)
}

func addAction(c *gin.Context) {
	addResource(&exercise.Action{}, c)
}

func getActions(c *gin.Context) {
	getResource(&exercise.Action{}, c)

}
