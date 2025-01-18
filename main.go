package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	
	"seqd/core"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(Help, version)

func main() {
	log.SetFlags(0)
	log.SetPrefix("[seqd] ")
	runtime.GOMAXPROCS(1)
	
	// Short CLI args parser
	args, err := core.ParseArgs(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	
	if args.Help {
		fmt.Println(help)
		os.Exit(0)
	}
	
	if args.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	
	err = core.PrintDateRange(os.Stdout, &args)
	if err != nil {
		log.Fatalln(err)
	}
}
