package pizza

import "time"

type pizzaRepository struct {
	menu map[string]Pizza
}

func NewPizzaRepository() Repository {
	return &pizzaRepository{
		map[string]Pizza{
			"Cheese": {Name: "Pizza Cheese", Duration: 3 * time.Second},
			"BBQ":    {Name: "Pizza BBQ", Duration: 5 * time.Second},
		},
	}
}

type Repository interface {
	GetMenuByKey(req string) *Pizza
	GetAllMenu() []Pizza
}

func (p *pizzaRepository) GetMenuByKey(req string) *Pizza {
	pizza, found := p.menu[req]
	if !found {
		return nil
	}
	return &pizza
}

func (p *pizzaRepository) GetAllMenu() []Pizza {
	var results []Pizza
	for _, val := range p.menu {
		results = append(results, val)
	}
	return results
}
