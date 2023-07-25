package order

import (
	"log"
	"pizza-hub/common"
	"pizza-hub/domain/chef"
	"pizza-hub/domain/pizza"
	"sync"
	"time"
)

var (
	chefsMutex sync.Mutex

	chefPool = sync.Pool{
		New: func() interface{} {
			return &chef.Chef{}
		},
	}
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

func (o *orderImplementation) PlaceOrder(pizzaType string) error {
	p := o.pizzaRepo.GetMenuByKey(pizzaType)
	if p == nil {
		return common.ErrSelectedMenuDoesntExist
	}

	chefs := o.chefRepo.ListChefs()
	totalChefs := len(chefs)
	if totalChefs == 0 {
		return common.ErrNoAvailableChefs
	}

	go func() {
		newOrder := new(Order)
		newOrder.PizzaType = pizzaType
		newOrder.Duration = p.Duration

		availableChefs := make(chan *chef.Chef, len(chefs))
		for _, chef := range chefs {
			availableChefs <- chef
		}

		// Process orders using available chefs concurrently
		for {
			select {
			case chef := <-availableChefs:
				go func() {
					o.processOrder(chef, newOrder)
					availableChefs <- chef // Return the chef to the pool
				}()
			}
		}
	}()

	return nil
}

func (o *orderImplementation) processOrder(chef *chef.Chef, order *Order) {
	chef.Lock()
	defer func() {
		chef.Unlock() // Unlock the chef after processing the order
		log.Println("chef is done cooking ")
	}()
	log.Println("chef is cooking ", order.PizzaType)

	time.Sleep(order.Duration)
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
