package pizza

type PizzaImplementation struct {
	Repo Repository
}

func NewPizzaService(repo Repository) Service {
	return &PizzaImplementation{
		Repo: repo,
	}
}

type Service interface {
	GetMenuByKey(req string) *Pizza
	GetAllMenu() []Pizza
}

func (p *PizzaImplementation) GetMenuByKey(req string) *Pizza {
	return p.Repo.GetMenuByKey(req)
}

func (p *PizzaImplementation) GetAllMenu() []Pizza {
	return p.Repo.GetAllMenu()
}
