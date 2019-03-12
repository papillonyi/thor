package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db, _ = gorm.Open("mysql", "root:root@tcp(mysql-1:3306)/thor?charset=utf8mb4&parseTime=True&loc=Local")

func Migrate() {
	db.LogMode(true)
	db.AutoMigrate(&CurrencyType{}, &CurrencyRate{}, &MyCurrencyData{}, &Task{})

}
