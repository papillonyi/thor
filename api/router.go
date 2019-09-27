package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/middleware/jwt"
	"github.com/papillonyi/thor/pkg/setting"
)

func InitRouter() *gin.Engine {

	store := cookie.NewStore([]byte("secret"))

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://loki.qijiucao.top"}
	config.AllowCredentials = true

	router.Use(gin.Recovery(), cors.New(config), sessions.Sessions("mysession", store))

	gin.SetMode(setting.ServerSetting.RunMode)

	router.GET("/auth", GetAuth)

	currency := router.Group("/api/v1/currency")
	currency.Use(jwt.JWT())
	{
		currency.POST("/exchange-rate/update", updateExchangeRate)
		currency.GET("/test/time", testTime)
		currency.GET("/scur/:scur/tcur/:tcur", getExchangeRate)
		currency.GET("/currency-type/", getAllCurrencyType)
	}

	trainer := router.Group("/api/v1/trainer")
	{
		trainer.POST("/post", getTrainer)
		trainer.GET("/get", getTrainer)
		trainer.GET("/gettest", getTest)
	}

	exercise := router.Group("/api/v1/exercise")
	{
		exercise.POST("/post", addAction)
		exercise.GET("/getActions", getActions)
		exercise.GET("/gettest", getTest)
	}

	//router.Run(":8080")
	return router
}
