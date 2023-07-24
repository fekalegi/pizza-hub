package order

import "time"

type Order struct {
	PizzaType string
	Duration  time.Duration
}
