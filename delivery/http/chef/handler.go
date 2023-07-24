package chef

import (
	"encoding/json"
	"net/http"
	"pizza-hub/common"
	"pizza-hub/delivery/http/chef/model"
	chefDomain "pizza-hub/domain/chef"
)

func (c *controller) addChef(w http.ResponseWriter, r *http.Request) {
	chef := new(chefDomain.Chef)
	total := c.chefService.Add(chef)
	resp := new(model.AddChefResponse)
	resp.TotalChefs = total
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.SuccessResponseWithData(resp, "success"))
}
