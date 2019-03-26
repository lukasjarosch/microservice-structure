package service

import (
	"io/ioutil"
	"testing"

	config2 "github.com/lukasjarosch/microservice-structure/internal/config"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var greetingTable = []struct {
	name        string
	greeting    string
	err         error
	expectError bool
}{
	{"Hans", "Hey there, Hans", nil, false},
	{"-1", "Hey there, -1", nil, false},
	{"", "", ErrEmptyName, true},
}

var farewellTable = []struct {
	name        string
	greeting    string
	err         error
	expectError bool
}{
	{"Hans", "See you soon, Hans", nil, false},
	{"-1", "See you soon, -1", nil, false},
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

		greeting, err := svc.Greeting(tt.name)

		if tt.expectError {
			g.Expect(err).To(HaveOccurred())
		} else {
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(greeting).To(BeEquivalentTo(tt.greeting))
		}
	}
}

// TestFarewell is a basic table-driven unit-test for the Farewell() RPC
func TestFarewell(t *testing.T) {
	g := NewGomegaWithT(t)

	for _, tt := range farewellTable {
		nopLog := logrus.New()
		nopLog.Out = ioutil.Discard

		config := config2.NewConfig()
		svc := NewExampleService(config, nopLog)

		greeting, err := svc.Farewell(tt.name)

		if tt.expectError {
			g.Expect(err).To(HaveOccurred())
		} else {
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(greeting).To(BeEquivalentTo(tt.greeting))
		}
	}
}
