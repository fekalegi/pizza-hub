package chef

type ChefImplementation struct {
	Repo        Repository
	ChefChannel chan *Chef
}

func NewChefService(repo Repository, chefChan chan *Chef) Service {
	return &ChefImplementation{
		Repo:        repo,
		ChefChannel: chefChan,
	}
}

type Service interface {
	Add(req *Chef) int
}

func (c *chefImplementation) Add(req *Chef) int {
	return c.repo.Add(req)
}
