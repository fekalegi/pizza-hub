package mocks

import (
	"pizza-hub/domain/pizza"
	"time"
)

// MockPizzaRepository is a mock implementation of the Repository interface for testing purposes.
type MockPizzaRepository struct{}

func (m *MockPizzaRepository) GetMenuByKey(key string) *pizza.Pizza {
	return &pizza.Pizza{Name: "Cheese", Duration: 3 * time.Second}
}

func (m *MockPizzaRepository) GetAllMenu() []pizza.Pizza {
	return []pizza.Pizza{
		{Name: "Cheese", Duration: 3 * time.Second},
		{Name: "BBQ", Duration: 5 * time.Second},
	}
}

// MockPizzaEmptyRepository is a mock implementation of the Repository interface with an empty menu for testing GetAllMenu().
type MockPizzaEmptyRepository struct{}

func (m *MockPizzaEmptyRepository) GetMenuByKey(key string) *pizza.Pizza {
	return nil
}

func (m *MockPizzaEmptyRepository) GetAllMenu() []pizza.Pizza {
	return nil
}
