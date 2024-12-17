package CalcServer

import "errors"

var (
	ErrInvalidExpression = errors.New("неправильное выражение")
	ErrDivisionByZero    = errors.New("деление на ноль")
	ErrEmptyExpression   = errors.New("пустая строка")
)