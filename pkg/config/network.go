package config

import "fmt"

// Network defines a common host:port configuration
type Network struct {
	Port int
	Host string
}

// Address returns the concatenated Host:Port combination as string
func (n *Network) Address() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}
