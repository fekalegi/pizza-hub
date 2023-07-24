package pizza

type pizzaImplementation struct {
	repo Repository
}

func NewPizzaService(repo Repository) Service {
	return &pizzaImplementation{
		repo: repo,
	}
}

type Service interface {
	GetMenuByKey(req string) *Pizza
	GetAllMenu() []Pizza
}

func (p *pizzaImplementation) GetMenuByKey(req string) *Pizza {
	return p.repo.GetMenuByKey(req)
}

func (p *pizzaImplementation) GetAllMenu() []Pizza {
	return p.repo.GetAllMenu()
}
