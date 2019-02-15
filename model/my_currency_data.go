package model

type MyCurrencyData struct {
	BaseMode
	Currency   CurrencyType
	CurrencyID uint
	Amount     float64
}
