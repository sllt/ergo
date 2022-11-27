package main

import (
	"flag"
	"fmt"

	"github.com/sllt/ergo"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/ergo/node"
)

func main() {
	flag.Parse()

	fmt.Println("")
	fmt.Println("to stop press Ctrl-C")
	fmt.Println("")

	apps := []gen.ApplicationBehavior{
		createDemoApp(),
	}
	opts := node.Options{
		Applications: apps,
	}
	demoNode, err := ergo.StartNode("app@localhost", "cookie", opts)
	if err != nil {
		panic(err)
	}
	demoNode.Wait()
}
