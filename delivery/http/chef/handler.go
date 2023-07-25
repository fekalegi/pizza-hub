package chef

import (
	"encoding/json"
	"net/http"
	"pizza-hub/common"
	"pizza-hub/delivery/http/chef/model"
)

func (c *controller) addChef(w http.ResponseWriter, r *http.Request) {
	chef, total := c.chefService.Add()
	response := model.AddChefResponse{
		ID:         chef.ID,
		TotalChefs: total,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.SuccessResponseWithData(response, "success"))
}
