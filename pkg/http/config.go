package http

import "fmt"

type Network struct {
	Port int
	Host string
}

// Address returns the concatenated Host:Port combination as string
func (n *Network) Address() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}
