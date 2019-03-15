package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/papillonyi/thor/pkg/setting"
	"log"
)

var db *gorm.DB

//var db, _ = gorm.Open("mysql", "root:root@tcp(mysql-1:3306)/thor?charset=utf8mb4&parseTime=True&loc=Local")

//func Migrate() {
//	db.LogMode(true)
//	db.AutoMigrate(&CurrencyType{}, &CurrencyRate{}, &MyCurrencyData{}, &Task{})
//
//}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	//db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	db.AutoMigrate(&CurrencyType{}, &CurrencyRate{}, &MyCurrencyData{}, &Task{}, &Auth{})
}
