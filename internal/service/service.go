package service

import (
	"fmt"

	"errors"

	"github.com/lukasjarosch/microservice-structure/internal/config"
	"github.com/sirupsen/logrus"
)

// ExampleService is the actual business-logic which you want to provide
type ExampleService struct {
	config *config.Config
	logger *logrus.Logger
}

var (
	ErrEmptyName = errors.New("the given name is empty")
)

// NewExampleService returns our business-implementation of the ExampleService
func NewExampleService(config *config.Config, logger *logrus.Logger) *ExampleService {

	service := &ExampleService{
		logger: logger,
		config: config,
	}

	return service
}

// Greeting implements the business-logic for this RPC
func (e *ExampleService) Greeting(name string) (greeting string, err error) {
	if name == "" {
		return "", ErrEmptyName
	}

	return fmt.Sprintf("Hey there, " + name), nil
}

// Farewell implements the business-logic for this RPC
func (e *ExampleService) Farewell(name string) (farewell string, err error) {
	if name == "" {
		return "", ErrEmptyName
	}

	return fmt.Sprintf("See you soon, " + name), nil
}
