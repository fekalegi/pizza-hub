package order

import (
	"net/http"
	orderDomain "pizza-hub/domain/order"
)

type controller struct {
	orderService orderDomain.Service
}

// NewOrderController : Instance for register Order Service
func NewOrderController(orderService orderDomain.Service) *controller {
	return &controller{orderService: orderService}
}

func (c *controller) RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.placeOrder(w, r)
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}
