package main

import (
	"github.com/christopherkeim/go-web-template/internal/server"
)

func main() {
	server.RunHTTPServerOnAddress("0.0.0.0:8000")
}
