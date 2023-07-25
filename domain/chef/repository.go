package chef

import (
	"math/rand"
)

type repository struct {
	chefs   []*Chef
	chefsMu *sync.RWMutex
}

func NewChefRepository() Repository {
	return &repository{
		make([]*Chef, 0),
		&sync.RWMutex{},
	}
}

type Repository interface {
	Add(req *Chef) int
	ListChefs() []*Chef
	GetChefsCount() int
}

func (r *repository) Add(req *Chef) int {
	r.chefsMu.Lock()
	defer r.chefsMu.Unlock()

	req.ID = len(r.chefs) + 1
	r.chefs = append(r.chefs, req)
	return req, len(r.chefs)
}

func (r *repository) ListChefs() []*Chef {
	r.chefsMu.RLock()
	defer r.chefsMu.RUnlock()
	return r.chefs
}

func (r *repository) GetChefsCount() int {
	r.chefsMu.RLock()
	chefsCount := len(r.chefs)
	defer r.chefsMu.RUnlock()
	return chefsCount
}
