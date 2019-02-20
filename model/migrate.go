package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db, _ = gorm.Open("mysql", "root:root@/thor?charset=utf8mb4&parseTime=True&loc=Local")

func Migrate() {
	//db, err := gorm.Open("mysql", "root:root@/thor?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&CurrencyType{}, &CurrencyRate{}, &MyCurrencyData{})
	db.Model(&CurrencyRate{}).AddForeignKey("source_currency_id", "currency_types(id)", "CASCADE", "RESTRICT")
	db.Model(&CurrencyRate{}).AddForeignKey("to_currency_id", "currency_types(id)", "CASCADE", "RESTRICT")
	db.Model(&MyCurrencyData{}).AddForeignKey("currency_id", "currency_types(id)", "CASCADE", "RESTRICT")

	//UpdateAllExchangeRate(db)
}
