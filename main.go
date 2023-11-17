package main

import (
	"fmt"

	"github.com/aristanetworks/goeapi"
)

type Conn struct {
	Transport string
	Host      string
	Username  string
	Password  string
	Port      int
	Config    string
}

type VersionResp struct {
	ModelName        string  `json:"modelName"`
	InternalVersion  string  `json:"internalVersion"`
	SystemMacAddress string  `json:"systemMacAddress"`
	SerialNumber     string  `json:"serialNumber"`
	MemTotal         int     `json:"memTotal"`
	BootupTimestamp  float64 `json:"bootupTimestamp"`
	MemFree          int     `json:"memFree"`
	Version          string  `json:"version"`
	Architecture     string  `json:"architecture"`
	InternalBuildID  string  `json:"internalBuildId"`
	HardwareRevision string  `json:"hardwareRevision,omitempty"`
}

func (s *VersionResp) GetCmd() string {
	return "show version"
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
		Username:  "admin",
		Password:  "admin",
		Port:      80,
	}
	// Use the connection method to connect to the device.
	Connect, err := d.Connect()
	if err != nil {
		fmt.Println(err)
	}
	// Print the running-config as a massive string
	RunningConfig := Connect.RunningConfig()
	fmt.Println(RunningConfig + "\n")

	// Run some regular commands get the map[string]string output
	fmt.Println("Running a show version \n")
	commands := []string{"show version"}
	conf, err := Connect.Enable(commands)
	if err != nil {
		panic(err)
	}
	for k, v := range conf[0] {
		fmt.Println(k, v)
	}
	fmt.Print(conf[0])

	// Run some regular commands get the response typed the way we want it.

	Showversion := &VersionResp{}
	handle, err := Connect.GetHandle("json")
	if err != nil {
		fmt.Println(err)
	}
	handle.AddCommand(Showversion)
	if err := handle.Call(); err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	fmt.Printf("Version           : %s\n", Showversion.Version)
	fmt.Printf("System MAC        : %s\n", Showversion.SystemMacAddress)
	fmt.Printf("Serial Number     : %s\n", Showversion.SerialNumber)
}
