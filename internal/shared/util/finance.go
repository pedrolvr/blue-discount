package util

import (
	"github.com/shopspring/decimal"
)

const CentsFactor = 100 // 2 decimal places

func CentsToDecimal(val int64) decimal.Decimal {
	value := decimal.NewFromInt(val)
	centsFactor := decimal.NewFromInt(CentsFactor)
	return value.Div(centsFactor)
}

func DecimalToCents(val decimal.Decimal) int64 {
	centsFactor := decimal.NewFromInt(CentsFactor)
	return val.Mul(centsFactor).BigInt().Int64()
}

func DiscountInDecimal(value decimal.Decimal, perc int32) decimal.Decimal {
	percent := decimal.NewFromInt32(perc)
	return value.Abs().Mul(percent.Div(decimal.NewFromInt(100)))
}

func DiscountInCents(val int64, perc int32) int64 {
	return DecimalToCents(DiscountInDecimal(CentsToDecimal(val), perc))
}
