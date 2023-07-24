package pizza

import (
	"encoding/json"
	"net/http"
	"pizza-hub/common"
	"pizza-hub/delivery/http/pizza/model"
	domain "pizza-hub/domain/pizza"
)

func (c *controller) getListMenus(w http.ResponseWriter, r *http.Request) {
	pizzas := c.pizzaService.GetAllMenu()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.SuccessResponseWithData(mapPizzaResponse(pizzas), "success"))
}

func mapPizzaResponse(p []domain.Pizza) []model.PizzaMenusResponse {
	var res []model.PizzaMenusResponse
	for _, val := range p {
		res = append(res, model.PizzaMenusResponse{
			Name:     val.Name,
			Duration: val.Duration.String(),
		})
	}

	return res
}
