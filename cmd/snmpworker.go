package main

import (
	"log"

	"github.com/itiic/snmpworker/pkg/async"
	"github.com/itiic/snmpworker/pkg/conf"
	flag "github.com/spf13/pflag"
)

var nodes string
var config string

func init() {
	flag.StringVar(&nodes, "nodes", "nodes", "Nodes")
	flag.StringVar(&config, "config", "config.json", "Config")
	flag.Parse()
}

func main() {

	// channels
	inChan := make(chan string)
	outChan := make(chan string)

	// load config
	cfg, err := conf.NewConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Load data
	go async.Load(nodes, inChan)

	// Load balance workload
	go async.FanOutFanIn(inChan, outChan, cfg)

	for v := range outChan {
		log.Println(v)
	}
}
