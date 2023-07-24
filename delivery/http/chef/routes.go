package chef

import (
	"net/http"
	chefDomain "pizza-hub/domain/chef"
)

type controller struct {
	chefService chefDomain.Service
}

// NewChefController : Instance for register Chef Service
func NewChefController(chefService chefDomain.Service) *controller {
	return &controller{chefService: chefService}
}

func (c *controller) RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.addChef(w, r)
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}
