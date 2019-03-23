package service

import (
	"context"
	"testing"

	"io/ioutil"

	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
	"github.com/lukasjarosch/microservice-structure/internal/config"
	"github.com/lukasjarosch/microservice-structure/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	mockRepository := database.MockRepository{}
	nopLogger := logrus.New()
	nopLogger.Out = ioutil.Discard
	cfg := &config.Config{}

	tests := []struct {
		name string
		greeting *greeter.GreetingResponse
		err error
	} {
		{
			"Hans",
			&greeter.GreetingResponse{
				Greeting: "Greetings, Hans",
			},
			nil,
		},
		{
			"-1",
			&greeter.GreetingResponse{
				Greeting: "Greetings, -1",
			},
			nil,
		},
		{
			"Peter",
			&greeter.GreetingResponse{
				Greeting: "Greetings, Peter",
			},
			nil,
		},
		{
			"",
			nil,
			ErrNameEmpty,
		},
		{
			"Donald Trump",
			nil,
			ErrNameForbidden,
		},
	}

	svc := exampleService{cfg, nopLogger, mockRepository}

	for _, tt := range tests {
		resp, err := svc.Greeting(context.Background(), &greeter.GreetingRequest{
			Name: tt.name,
		})

		assert.Equal(t, tt.greeting, resp)
		assert.Equal(t, tt.err, err)
	}

}
