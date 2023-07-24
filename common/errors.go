package common

import "errors"

var (
	ErrSelectedMenuDoesntExist = errors.New("the selected menu is not on our menu")
	ErrNoAvailableChefs        = errors.New("no chefs available to process the order")
)
