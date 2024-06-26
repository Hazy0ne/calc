package variables

var Roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}
var ConvIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var AIn, BIn *int
var Operators = map[string]func() int{
	"+": func() int { return *AIn + *BIn },
	"-": func() int { return *AIn - *BIn },
	"/": func() int { return *AIn / *BIn },
	"*": func() int { return *AIn * *BIn },
}
var Data []string

const (
	NoOperation = "Выдача паники, так как строка не является математической операцией."
	Long        = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"
	Different   = "Выдача паники, так как используются одновременно разные системы счисления."
	Negative    = "Выдача паники, так как в римской системе нет отрицательных чисел."
	Zero        = "Вывод паники, так как в римской системе нет числа меньше I."
	NoValid     = "Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами. От 1 до 10. "
)
