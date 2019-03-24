package service

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	config2 "github.com/lukasjarosch/microservice-structure/internal/config"
	"context"
	"github.com/lukasjarosch/microservice-structure-protobuf/greeter"
)

var greetingTable = []struct{
	name string
	greeting string
	err error
	expectError bool
}{
	{"Hans", "Hey there, Hans", nil, false},
	{"-1", "Hey there, -1", nil, false},
	{"", "", ErrEmptyName, true},
}

// TestGreeting is a basic table-driven unit-test for the Greeting() RPC
func TestGreeting(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, tt := range greetingTable {
		nopLog := logrus.New()
		nopLog.Out = ioutil.Discard

		config := config2.NewConfig()

		svc := NewExampleService(config, nopLog)

		response, err := svc.Greeting(context.Background(), &greeter.GreetingRequest{
			Name: tt.name,
		})

		if tt.expectError {
			g.Expect(err).To(HaveOccurred())
		} else {
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(response.Greeting).To(BeEquivalentTo(tt.greeting))
		}
	}
}
