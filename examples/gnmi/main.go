package main

import (
	"context"

	"github.com/aristanetworks/glog"
	"github.com/aristanetworks/goarista/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// Create a gnmi structure of the client data
var cfg = &gnmi.Config{
	Addr:     "172.20.20.2:6030",
	Username: "admin",
	Password: "admin",
}

func main() {
	// Use the system state path
	paths := []string{"/system/openconfig-system:state/"}
	// Origin of openconfig
	var origin = "openconfig"
	//Create a new context dialer passing in the config
	ctx := gnmi.NewContext(context.Background(), cfg)
	// Dial via grpc
	client, err := gnmi.Dial(cfg)
	if err != nil {
		glog.Fatal(err)
	}
	//Usr the NewGetRequest method and pass in some parameters
	req, err := gnmi.NewGetRequest(gnmi.SplitPaths(paths), origin)
	if err != nil {
		glog.Fatal(err)
	}
	if cfg.Addr != "" {
		if req.Prefix == nil {
			req.Prefix = &pb.Path{}
		}
		req.Prefix.Target = cfg.Addr
	}
	// This will initiate the GetWithRequest method.  It has a print statement within the method hence why we are not printing.
	err = gnmi.GetWithRequest(ctx, client, req)
	if err != nil {
		glog.Fatal(err)
	}
}
