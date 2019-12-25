package utils

import (
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

func AmountWalletRuleSave(value float64) int64 {
	d := decimal.NewFromFloat(value)
	dWithPlaces, _ := d.Round(2).Float64()
	return cast.ToInt64(dWithPlaces * 100)
}

func AmountWalletRuleExtract(value int64) float64 {
	valueToFloat := cast.ToFloat64(value)
	floatValue := valueToFloat / 100
	needRound := decimal.NewFromFloat(floatValue)
	result, _ := needRound.Round(2).Float64()
	return result
}
