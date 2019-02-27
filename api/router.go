package api

import (
	"github.com/gin-gonic/gin"
)

func Load() {
	router := gin.Default()
	router.Use(gin.Recovery())

	currency := router.Group("/api/v1/currency")
	{
		currency.POST("/exchange-rate/update", updateExchangeRate)
		currency.GET("/test/time", testTime)
		currency.GET("/scur/:scur/tcur/:tcur", getExchangeRate)
		currency.GET("/currency-type/", getAllCurrencyType)
	}
	router.Run(":8080")
}
