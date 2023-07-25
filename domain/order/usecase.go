package order

import (
	"log"
	"pizza-hub/common"
	"pizza-hub/domain/chef"
	"pizza-hub/domain/pizza"
	"sync"
	"time"
)

type OrderImplementation struct {
	ChefRepo    chef.Repository
	PizzaRepo   pizza.Repository
	ChefChannel chan *chef.Chef
}

func NewOrderService(chefRepo chef.Repository, pizzaRepo pizza.Repository, chefChannel chan *chef.Chef) Service {
	return &OrderImplementation{
		ChefRepo:    chefRepo,
		PizzaRepo:   pizzaRepo,
		ChefChannel: chefChannel,
	}
}

type Service interface {
	PlaceOrder(pizzaType string) error
}

func (o *OrderImplementation) PlaceOrder(pizzaType string) error {
	p := o.PizzaRepo.GetMenuByKey(pizzaType)
	if p == nil {
		return common.ErrSelectedMenuDoesntExist
	}

	chefs := o.ChefRepo.ListChefs()
	totalChefs := len(chefs)
	if totalChefs == 0 {
		return common.ErrNoAvailableChefs
	}

	availChef := <-o.ChefChannel
	newOrders := &Order{
		PizzaType: pizzaType,
		Duration:  p.Duration,
	}

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		o.processOrder(availChef, newOrders, &wg)
	}()

	wg.Wait()

	return nil
}

func (o *OrderImplementation) processOrder(chef *chef.Chef, order *Order, wg *sync.WaitGroup) {
	defer func() {
		//chef.Unlock() // Unlock the chef after processing the order
		log.Println("chef is done cooking ", chef.ID)
		o.ChefChannel <- chef
		wg.Done()
	}()
	log.Println("chef", chef.ID, "is cooking ", order.PizzaType)

	time.Sleep(order.Duration)
}
