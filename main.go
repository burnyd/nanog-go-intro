package main

import (
	"fmt"

	"github.com/aristanetworks/goeapi"
)
// Structures the data in the way we want to srtucture it.
type Conn struct {
	Transport string
	Host      string
	Username  string
	Password  string
	Port      int
	Config    string
}
// Method returns a pointer to goeapi.Node and a error but connects to the device.
func (c *Conn) Connect() (*goeapi.Node, error) {
	connect, err := goeapi.Connect(c.Transport, c.Host, c.Username, c.Password, c.Port)
	if err != nil {
		fmt.Println(err)
	}
	return connect, nil
}

func main() {
	// Structure the connection data the way we want to strucgture it.
	d := Conn{
		Transport: "http",
		Host:      "172.20.20.2",
		Username:  "Admin",
		Password:  "Admin",
		Port:      80,
	}
	// Use the connection method to connect to the device.
	Connect, err := d.Connect()
	if err != nil {
		fmt.Println(err)
	}
	// Print the running-config as a massive string 
	RunningConfig := Connect.RunningConfig()
	fmt.Println(RunningConfig)

	// Run some regular commands get the map[string]string output
	// Run some regular commands get the response typed the way we want it. 

}