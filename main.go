package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"seqd/seqd"
)

var env = "development"
var version = "0.0.0"
var help = fmt.Sprintf(Help, version)

func main() {
	log.SetFlags(0)
	log.SetPrefix("seqd: ")
	runtime.GOMAXPROCS(1)

	args, err := seqd.ParseArgs(os.Args)
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

	err = seqd.GenerateDateRange(seqd.StdoutWriter{}, &args)
	if err != nil {
		log.Fatalln(err)
	}
}
