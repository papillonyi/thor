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
	router.Use(gin.Recovery(), cors.Default(), sessions.Sessions("mysession", store))
	gin.SetMode(setting.ServerSetting.RunMode)

	router.GET("/auth", GetAuth)

	currency := router.Group("/api/v1/currency")
	currency.Use(jwt.JWT(), cors.Default())
	{
		currency.POST("/exchange-rate/update", updateExchangeRate)
		currency.GET("/test/time", testTime)
		currency.GET("/scur/:scur/tcur/:tcur", getExchangeRate)
		currency.GET("/currency-type/", getAllCurrencyType)
	}
	//router.Run(":8080")
	return router
}
