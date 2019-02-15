package model

import "time"

type CurrencyRate struct {
	BaseMode
	SourceCurrency   CurrencyType
	SourceCurrencyID uint
	ToCurrency       CurrencyType
	ToCurrencyID     uint
	Rate             float64
	date             time.Time
}
