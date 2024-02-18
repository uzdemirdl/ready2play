package money

import (
	"github.com/shopspring/decimal"
)

type Money = decimal.Decimal

const (
	Precision       = 2
	kopecksInRouble = 100
)

func Zero() Money {
	return New(0.0)
}

func New(value float64) Money {
	return decimal.NewFromFloat(value).Round(Precision)
}

func FromKopecks(kopecks uint64) Money {
	return Divide(decimal.NewFromInt(int64(kopecks)), kopecksInRouble)
}

func Multiply(money Money, value float64) Money {
	decimalValue := decimal.NewFromFloat(value).Round(Precision)
	return money.Mul(decimalValue).Round(Precision)
}

func Divide(money Money, value int) Money {
	decimalValue := decimal.NewFromInt(int64(value))
	return money.Div(decimalValue).Round(Precision)
}
