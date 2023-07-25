package chef_test

import (
	"pizza-hub/domain/chef"
	"pizza-hub/mocks"
	"reflect"
	"testing"
)

func Test_chefImplementation_Add(t *testing.T) {
	type fields struct {
		repo        chef.Repository
		chefChannel chan *chef.Chef
	}
	tests := []struct {
		name   string
		fields fields
		want   *chef.Chef
		want1  int
	}{
		{
			name:   "Success",
			fields: fields{repo: &mocks.MockChefRepository{}},
			want:   &chef.Chef{ID: 1},
			want1:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &chef.ChefImplementation{
				Repo:        tt.fields.repo,
				ChefChannel: tt.fields.chefChannel,
			}
			got, got1 := c.Add()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Add() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
