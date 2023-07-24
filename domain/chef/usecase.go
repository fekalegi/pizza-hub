package chef

type chefImplementation struct {
	repo Repository
}

func NewChefService(repo Repository) Service {
	return &chefImplementation{
		repo: repo,
	}
}

type Service interface {
	Add(req *Chef) int
}

func (c *chefImplementation) Add(req *Chef) int {
	return c.repo.Add(req)
}
