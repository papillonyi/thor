package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type CurrencyRate struct {
	BaseMode
	SourceCurrency   CurrencyType
	SourceCurrencyID uint
	ToCurrency       CurrencyType
	ToCurrencyID     uint
	Rate             float64
	date             time.Time
}

//type resultType struct {
//	status	string
//	scur	string
//	tcur	string
//	ratenm	string
//	rate 	float32
//	update	time.Time
//}
//
//type financeRateType struct {
//	success	int
//	result	resultType
//}

func getExchangeRate(scur string, tcur string) (float64, time.Time) {
	info := getFinanceInfo(scur, tcur)
	//fmt.Println(info)
	m := info.(map[string]interface{})
	result := m["result"].(map[string]interface{})

	rate, err := strconv.ParseFloat(result["rate"].(string), 64)
	if err != nil {
		panic(err)
	}

	str := "2006-01-02 15:04:05"

	date, err := time.Parse(str, result["update"].(string))
	if err != nil {
		panic(err)
	}

	return rate, date
}

func getFinanceInfo(scur string, tcur string) interface{} {
	url := fmt.Sprintf("http://api.k780.com/?app=finance.rate&scur=%s&tcur=%s&appkey=34013&sign=d0f7990ff294d38f3383768c49effa86&format=json", scur, tcur)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	var target interface{}

	json.NewDecoder(res.Body).Decode(&target)

	//fmt.Println(res)
	return target
}

func updateExchangeRate(scur CurrencyType, tcur CurrencyType, db *gorm.DB) {
	fmt.Println(scur.CurrencyName, tcur.CurrencyName)
	rate, date := getExchangeRate(scur.CurrencyName, tcur.CurrencyName)
	currencyRate := CurrencyRate{
		SourceCurrencyID: scur.ID,
		ToCurrencyID:     tcur.ID,
		Rate:             rate,
		date:             date,
	}
	fmt.Println(currencyRate)
	db.Create(&currencyRate)
}

func UpdateAllExchangeRate(db *gorm.DB) {
	var currencyTypes []CurrencyType
	db.Find(&currencyTypes)
	db.Select("id, currency_name").Find(&currencyTypes)
	fmt.Println(currencyTypes)
	for _, scur := range currencyTypes {
		scurName := scur.CurrencyName
		for _, tcur := range currencyTypes {
			tcurName := tcur.CurrencyName
			if tcurName != scurName {
				updateExchangeRate(scur, tcur, db)
				time.Sleep(2 * time.Second)
			}

		}

	}

}
