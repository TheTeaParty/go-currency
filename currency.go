package currency

type Currency struct {
	Code     string
	Fraction int
}

var currencies = map[string]*Currency{
	"AMD":           {Code: "AMD", Fraction: 2},
	"USD":           {Code: "USD", Fraction: 2},
	"RUB":           {Code: "RUB", Fraction: 2},
	"TOKENIZER-FUN": {Code: "TOKENIZER-FUN", Fraction: 2},
	"TOKENIZER":     {Code: "TOKENIZER", Fraction: 2},
}
