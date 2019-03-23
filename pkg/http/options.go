package http

// Option defines a settable option of the HTTP server
type Option func(*Options)

// Options defines all settable options of the HTTP server
type Options struct {
	Network    Network
	GRPCServer Network
}

