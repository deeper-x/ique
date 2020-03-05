// Package network is delegated to network properties and actions
package network

import (
	"log"
	"net"
	"time"
)

// Checker is the network interface
type Checker interface {
	IsRunning() bool
}

// Service is the core instance
type Service struct {
	URI string
}

// IsRunning check if host:port is available
func (s Service) IsRunning() bool {
	_, err := net.DialTimeout("tcp", s.URI, time.Duration(1)*time.Second)

	if err != nil {
		log.Printf("Error in connection: %v", err)
		return false
	}

	return true
}
