package database

type Repository interface {
	GetGreeting(name string) (*Greeting, error)
	SetGreeting(greeting Greeting) error
}

// Greeting is the model for storing a greeting for a certain name
type Greeting struct {
	Name string
	Text string
}
