package builtin

import "errors"

var (
	ErrCannotAddNegativeItem    = errors.New("cannot add items with negative quantities")
	ErrItemNotExist             = errors.New("item not exist")
	ErrItemInsufficientQuantity = errors.New("item insufficient quantity")
	ErrItemContainerIsFull      = errors.New("item container is full")
)
