package pizza

import (
	"net/http"
	pizzaDomain "pizza-hub/domain/pizza"
)

type controller struct {
	pizzaService pizzaDomain.Service
}

// NewPizzaController : Instance for register Pizza Service
func NewPizzaController(pizzaService pizzaDomain.Service) *controller {
	return &controller{pizzaService: pizzaService}
}

func (c *controller) RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.getListMenus(w, r)
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}
