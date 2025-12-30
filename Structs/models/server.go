package models

import "fmt"

type server struct {
	port int
}

func NewServer(port_number int) *server {

	if port_number <= 0 {
		panic("port_number_less_than_equal_to_0")
	}

	return &server{
		port: port_number,
	}
}

func (s *server) Start() {
	fmt.Printf("SERVER_STARTED %d\n", s.port)
}

func (s *server) Port() int {
	return s.port
}
