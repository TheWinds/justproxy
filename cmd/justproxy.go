package main

import (
	"github.com/thewinds/justproxy"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var configFileName = kingpin.Arg("conf", "proxy config file").Required().String()

func main() {
	kingpin.Parse()
	justproxy.Proxy(*configFileName)
}
