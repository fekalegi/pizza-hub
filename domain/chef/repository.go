package chef

import (
	"math/rand"
)

type repository struct {
	chefs []*Chef
}

func NewChefRepository() Repository {
	return &repository{
		make([]*Chef, 0),
	}
}

type Repository interface {
	Add() (*Chef, int)
	ListChefs() []*Chef
	GetChefsCount() int
}

func (r *repository) Add() (*Chef, int) {
	req := new(Chef)
	req.ID = rand.Int()
	r.chefs = append(r.chefs, req)
	return req, len(r.chefs)
}

func (r *repository) ListChefs() []*Chef {
	return r.chefs
}

func (r *repository) GetChefsCount() int {
	chefsCount := len(r.chefs)
	return chefsCount
}
