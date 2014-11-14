package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/libcluster"
	"github.com/docker/libcluster/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s node1 node2 ...\n", os.Args[0])
		os.Exit(1)
	}

	// Setup logging
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	c := libcluster.NewCluster()
	for _, addr := range os.Args[1:] {
		n := libcluster.NewNode(addr, addr)
		if err := n.Connect(nil); err != nil {
			log.Fatal(err)
		}
		if err := c.AddNode(n); err != nil {
			log.Fatal(err)
		}
	}
	log.Fatal(api.ListenAndServe(c, ":4243"))
}