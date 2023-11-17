// main_test.go
package main

import (
	"testing"

	"github.com/aristanetworks/goeapi"
)

// MockConn is a mock implementation of the Conn struct for testing
type MockConn struct {
	Transport string
	Host      string
	Username  string
	Password  string
	Port      int
}

// Mock the Connect method for testing
func (c *MockConn) Connect() (*goeapi.Node, error) {
	connect, err := goeapi.Connect(c.Transport, c.Host, c.Username, c.Password, c.Port)
	if err != nil {
		return nil, err
	}
	return connect, nil
}

func TestConnect(t *testing.T) {
	// Create a mock Conn with your test data
	mockConn := &MockConn{
		Transport: "http",
		Host:      "mockHost",
		Username:  "mockUsername",
		Password:  "mockPassword",
		Port:      1234,
	}

	// Call the Connect method on the mockConn
	node, err := mockConn.Connect()

	// Check if there was an error
	if err != nil {
		t.Errorf("Connect() returned an error: %v", err)
	}

	// Check if the returned Node is not nil
	if node == nil {
		t.Error("Connect() returned a nil Node")
	}

	// You can add more specific checks based on your requirements
	// For example, check if the values of the returned Node match your expectations
	if mockConn != nil && mockConn.Host != "mockHost" {
		t.Errorf("Unexpected mockConn.Host. Expected: %s, Got: %s", "mockHost", mockConn.Host)
	}
}
