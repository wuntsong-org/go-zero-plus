package main

import (
	"flag"

	"github.com/wuntsong-org/go-zero-plus/core/conf"
	"github.com/wuntsong-org/go-zero-plus/gateway"
)

var configFile = flag.String("f", "etc/gateway.yaml", "config file")

func main() {
	flag.Parse()

	var c gateway.GatewayConf
	conf.MustLoad(*configFile, &c)
	gw := gateway.MustNewServer(c)
	defer gw.Stop()
	gw.Start()
}
