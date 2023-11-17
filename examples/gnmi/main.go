package main

import (
	"context"

	"github.com/aristanetworks/glog"
	"github.com/aristanetworks/goarista/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

var cfg = &gnmi.Config{
	Addr:     "172.20.20.2:6030",
	Username: "admin",
	Password: "admin",
}

func main() {
	paths := []string{"/system/openconfig-system:state/"}
	var origin = "openconfig"
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		glog.Fatal(err)
	}

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

	err = gnmi.GetWithRequest(ctx, client, req)
	if err != nil {
		glog.Fatal(err)
	}
	//fmt.Println(err)
}
