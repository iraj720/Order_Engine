package utils

import (
	"github.com/shopspring/decimal"
)

func ConvertStringToDecimal(number string) decimal.Decimal {
	res, err := decimal.NewFromString(number)
	if err != nil {
		panic("cannot convert string to decimal")
	}
	return res
}

// func ConvertMarket(ms ...*orderbook.Market) []*order_engine.Market {

// }
