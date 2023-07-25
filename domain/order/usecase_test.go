package order_test

import (
	"errors"
	"pizza-hub/common"
	"pizza-hub/domain/chef"
	"pizza-hub/domain/order"
	"pizza-hub/domain/pizza"
	"pizza-hub/mocks"
	"testing"
)

func Test_orderImplementation_PlaceOrder(t *testing.T) {
	ch := make(chan *chef.Chef)
	type fields struct {
		chefRepo  chef.Repository
		pizzaRepo pizza.Repository
		chefChan  chan *chef.Chef
	}
	type args struct {
		pizzaType string
	}
	go func() {
		ch <- &chef.Chef{ID: 1}
	}()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Success",
			fields: fields{
				chefRepo:  &mocks.MockChefRepository{},
				pizzaRepo: &mocks.MockPizzaRepository{},
				chefChan:  ch,
			},
			args:    args{pizzaType: "Cheese"},
			wantErr: false,
			err:     nil,
		}, {
			name: "Pizza is not on the menu",
			fields: fields{
				chefRepo:  &mocks.MockChefEmptyRepository{},
				pizzaRepo: &mocks.MockPizzaEmptyRepository{},
			},
			args:    args{pizzaType: "Pepperoni"},
			wantErr: true,
			err:     common.ErrSelectedMenuDoesntExist,
		}, {
			name: "No available Chef",
			fields: fields{
				chefRepo:  &mocks.MockChefEmptyRepository{},
				pizzaRepo: &mocks.MockPizzaRepository{},
			},
			args:    args{pizzaType: "Pepperoni"},
			wantErr: true,
			err:     common.ErrNoAvailableChefs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &order.OrderImplementation{
				ChefRepo:    tt.fields.chefRepo,
				PizzaRepo:   tt.fields.pizzaRepo,
				ChefChannel: tt.fields.chefChan,
			}
			if err := o.PlaceOrder(tt.args.pizzaType); (err != nil && errors.Is(err, tt.err)) != tt.wantErr {
				t.Errorf("PlaceOrder() error = %v, wantErr %v err must equal to %v", tt.err, tt.wantErr, err)
			}
		})
	}
}
