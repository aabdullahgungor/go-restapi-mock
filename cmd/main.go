package main

import (
	"github.com/aabdullahgungor/go-restapi-mock/server"
)

func main() {
	

	s := server.NewServer()
	s.Run()
}