package model

import (
	"time"
)

type BaseMode struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CurrencyType struct {
	BaseMode
	CurrencyName string `gorm:"unique;not null;size:128"`
}

func GetAllCurrencyType() []CurrencyType {
	var vts []CurrencyType
	db.Find(&vts)
	return vts
}
