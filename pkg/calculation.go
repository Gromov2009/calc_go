package calculation

import (
	"strconv"
	"unicode"
)

type Result struct {
	result float64
	rest   string
	err    error
}

func Calc(exp string) (float64, error) {
	if exp == "" {
		return 0.0, ErrInvalidExpression //errors.New("empty body")
	}
	res := plusMinus(exp)
	return res.result, res.err
}

func Bracket(exp string) Result {

	if len(exp) == 0 {
		return Result{0, "", ErrInvalidExpression} // errors.New("error in expression")}
	}

	firstChar := rune(exp[0])

	if firstChar == rune('(') {
		res := plusMinus(string([]rune(exp)[1:]))
		if len(res.rest) > 0 && rune(res.rest[0]) == rune(')') {
			res.rest = string([]rune(res.rest)[1:])
		} else {
			return Result{0, "", ErrMissingClosingBracket}
		}
		return res
	}

	return extractNumber(exp)
}

func mulDiv(exp string) Result {
	current := Bracket(exp)
	if current.err != nil {
		return current
	}

	for len(current.rest) > 0 {

		if rune(current.rest[0]) != rune('*') && rune(current.rest[0]) != rune('/') {
			break
		}
		sign := rune(current.rest[0])

		nextExp := string([]rune(current.rest)[1:])

		result := current.result

		current = Bracket(nextExp)
		if current.err != nil {
			return current
		}

		if sign == rune('*') {
			result = result * current.result
		} else {
			if current.result == 0.0 {
				return Result{0, "", ErrDivByZero}
			} else {
				result = result / current.result
			}
		}

		current.result = result

	}

	return Result{current.result, current.rest, nil}

}

func plusMinus(exp string) Result {

	current := mulDiv(exp)
	if current.err != nil {
		return current
	}
	result := current.result

	for len(current.rest) > 0 {

		if rune(current.rest[0]) != rune('-') && rune(current.rest[0]) != rune('+') {
			break
		}
		sign := rune(current.rest[0])

		nextExp := string([]rune(current.rest)[1:])

		current = mulDiv(nextExp)
		if current.err != nil {
			return current
		}

		if sign == rune('-') {
			result = result - current.result
		} else {
			result = result + current.result
		}

		current.result = result

	}

	return Result{current.result, current.rest, nil}

}

// возвращает/проверяет цифры из переданной строки и остаток
func extractNumber(exp string) Result {

	var isNegative bool
	var j int

	if len(exp) > 0 {

		if rune(exp[0]) == rune('-') {
			isNegative = true
			exp = string([]rune(exp)[1:])
		}

		digits := []rune{}

		for _, char := range exp {
			if unicode.IsDigit(char) || char == '.' {
				digits = append(digits, char)
			} else {
				break
			}
		}

		j = len(digits)

		theRest := string([]rune(exp)[j:])

		digitsFromExp, err := strconv.ParseFloat(string(digits), 64)
		if err != nil {
			//log.Fatal(err)
			return Result{0, "", err}
		}

		if isNegative {
			return Result{-digitsFromExp, theRest, nil}
		} else {
			return Result{digitsFromExp, theRest, nil}
		}

	} else {
		return Result{0, "", ErrInvalidExpression}
	}

}
