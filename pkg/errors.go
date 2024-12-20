package calculation

import (
	"errors"
)

var (
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrMissingClosingBracket = errors.New("missing closing bracket")
	ErrDivByZero             = errors.New("division by zero")
)
