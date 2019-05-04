package currency

import "math"

type Money struct {
	Amount   uint64
	Currency *Currency
}

func (m *Money) GetFloat() float64 {
	return float64(m.Amount) / float64(10^m.Currency.Fraction)
}

func New(a uint64, c string) Money {

	currency := currencies[c]
	return Money{Amount: a, Currency: currency}
}

func NewFromFloat(a float64, c string) Money {

	currency := currencies[c]
	amount := math.Round(a*100) / 100 * float64(10^currency.Fraction)

	return Money{Amount: uint64(amount), Currency: currency}
}
