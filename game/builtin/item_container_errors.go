package builtin

import "errors"

var (
	ErrCannotAddNegativeOrZeroItem = errors.New("cannot add items with negative quantities or zero")
	ErrItemNotExist                = errors.New("item not exist")
	ErrItemInsufficientQuantity    = errors.New("item insufficient quantity")
	ErrItemContainerIsFull         = errors.New("item container is full")
	ErrItemContainerNotExist       = errors.New("item container not exist")
)
