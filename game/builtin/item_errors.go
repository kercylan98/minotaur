package builtin

import "errors"

var (
	ErrItemInsufficientQuantityDeduction = errors.New("insufficient quantity deduction")
	ErrItemStackLimit                    = errors.New("the number of items reaches the stacking limit")
)
