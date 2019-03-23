package database

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m MockRepository) GetGreeting(name string) (*Greeting, error) {
	args := m.Called(name)
	return args.Get(0).(*Greeting), args.Error(1)
}

func (m MockRepository) SetGreeting(greeting Greeting) error {
	args := m.Called(greeting)
	return args.Error(0)
}
