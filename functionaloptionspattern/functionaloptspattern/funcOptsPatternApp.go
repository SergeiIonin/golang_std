package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	option := WithPort(8080) // it's a func operating on options
	newOpts := options{}
	option(&newOpts)

	fmt.Println("New opts port", *newOpts.port)

	// using of Option is composable
	serverProps, err := ServerFullAddrWithOptions("localhost", WithPort(8888), WithTimeout(100))
	if err != nil {
		fmt.Errorf("Error occurred %w", err)
	}

	fmt.Println("server addr is ", serverProps.FullAddress)
	fmt.Println("server timeout is ", serverProps.Timeout)

	fmt.Println("---------")
	serverPropsEmpty, err := ServerFullAddrWithOptions("localhost")
	if err != nil {
		fmt.Errorf("Error occurred %w", err)
	}
	fmt.Println("server addr is ", serverPropsEmpty.FullAddress)
	fmt.Println("server timeout is ", serverPropsEmpty.Timeout)
}

type options struct {
	port    *Port
	timeout *int
}

type Option func(options *options) error

func WithPort(port Port) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("Port should be positive")
		}
		options.port = &port
		return nil
	}
}

func WithTimeout(timeout int) Option {
	return func(options *options) error {
		if timeout < 0 {
			return errors.New("Timeout should be positive")
		}
		options.timeout = &timeout
		return nil
	}
}

type Port int

type ServerProps struct {
	FullAddress string
	Timeout     int
}

func ServerFullAddrWithOptions(addr string, opts ...Option) (ServerProps, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return ServerProps{}, err
		}
	}

	var port Port
	if options.port == nil {
		port = defaultHttpPort()
	} else {
		if *options.port == 0 {
			port = randomHttpPort()
		} else {
			port = *options.port
		}
	}

	var timeout int
	if options.timeout == nil {
		timeout = 10
	} else {
		if *options.timeout == 0 {
			timeout = rand.Int()
		} else {
			timeout = *options.timeout
		}
	}

	return ServerProps{FullAddress: fmt.Sprintf("%s:%d", addr, port), Timeout: timeout}, nil
}

func defaultHttpPort() Port {
	return 8080
}

func randomHttpPort() Port {
	return 8888
}
