package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/papillonyi/thor/model"
	"strconv"
	"time"
)

func updateExchangeRate(c *gin.Context) {
	model.UpdateAllExchangeRate()
}

func testTime(c *gin.Context) {
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("10 second")
	}()
	fmt.Println("api done")
}

func getExchangeRate(c *gin.Context) {
	scur, _ := strconv.ParseUint(c.Param("scur"), 10, 64)
	tcur, _ := strconv.ParseUint(c.Param("tcur"), 10, 64)
	rate := model.GetRateByCurrencyAndDate(uint(scur), uint(tcur), time.Now())

	c.String(200, strconv.FormatFloat(rate, 'f', 6, 64))
}

func getAllCurrencyType(c *gin.Context) {
	vts := model.GetAllCurrencyType()
	c.JSON(200, vts)

}

func lengthOfLongestSubstring(str string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range str {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}
