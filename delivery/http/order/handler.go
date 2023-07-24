package order

import (
	"encoding/json"
	"errors"
	"net/http"
	"pizza-hub/common"
	"pizza-hub/delivery/http/order/model"
)

func (c *controller) placeOrder(w http.ResponseWriter, r *http.Request) {
	bodyRequest := new(model.PlaceOrderRequest)
	if err := json.NewDecoder(r.Body).Decode(&bodyRequest); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(common.BadRequestResponse(err))
		return
	}

	if bodyRequest.PizzaType == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(common.BadRequestResponse("please input pizza type"))
		return
	}

	err := c.orderService.PlaceOrder(bodyRequest.PizzaType)
	if err != nil && errors.Is(err, common.ErrSelectedMenuDoesntExist) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(common.BadRequestResponse(err.Error()))
		return
	} else if err != nil && errors.Is(err, common.ErrNoAvailableChefs) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(503)
		json.NewEncoder(w).Encode(common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(common.ErrorResponse(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(common.SuccessResponseNoData("success the order has been placed"))
}
