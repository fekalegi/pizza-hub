package main

import (
	"net/http"
	"pizza-hub/delivery/http/chef"
	"pizza-hub/delivery/http/order"
	"pizza-hub/delivery/http/pizza"
	chefDomain "pizza-hub/domain/chef"
	orderDomain "pizza-hub/domain/order"
	pizzaDomain "pizza-hub/domain/pizza"
)

func main() {

	chefRepo := chefDomain.NewChefRepository()
	newChefService := chefDomain.NewChefService(chefRepo)
	chefController := chef.NewChefController(newChefService)

	pizzaRepo := pizzaDomain.NewPizzaRepository()
	newPizzaService := pizzaDomain.NewPizzaService(pizzaRepo)
	pizzaController := pizza.NewPizzaController(newPizzaService)

	newOrderService := orderDomain.NewOrderService(chefRepo, pizzaRepo)
	orderController := order.NewOrderController(newOrderService)

	http.HandleFunc("/api/chef", chefController.RouteHandler)
	http.HandleFunc("/api/menus", pizzaController.RouteHandler)
	http.HandleFunc("/api/orders", orderController.RouteHandler)
	http.ListenAndServe(":8080", nil)
}
