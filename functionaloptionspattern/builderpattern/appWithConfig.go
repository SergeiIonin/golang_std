package main

import (
	"errors"
	"fmt"
)

func main() {
	config, err := ConfigBuilder{}.Port(8090).Build()
	if err != nil {
		fmt.Errorf("Error occurred %w", err)
	}
	fmt.Println("the port set up is ", config.Port)
}

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (cb ConfigBuilder) Port(port int) ConfigBuilder {
	cb.port = &port
	return cb
}

func (cb ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if cb.port == nil {
		cfg.Port = defaultHttpPort()
	} else {
		if *cb.port == 0 {
			cfg.Port = randomHttpPort()
		} else if *cb.port < 0 {
			return cfg, errors.New("Port number should be positive!")
		} else {
			cfg.Port = *cb.port
		}
	}
	return cfg, nil
}

func defaultHttpPort() int {
	return 8080
}

func randomHttpPort() int {
	return 8888
}

/*func NewServer(addr string, config Config) (*http.Server, error) {

}*/
