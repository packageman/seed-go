package main

import (
	goflag "flag"
	"log"

	flag "github.com/spf13/pflag"

	"github.com/packageman/seed/demo"
)

const (
	MODE_HELLO   = 1
	MODE_REQUEST = 2
	MODE_LOG     = 3
	MODE_REST    = 4
)

var (
	env  = flag.StringP("env", "e", "dev", "The running environment")
	mode = flag.IntP("mode", "m", 1, "Mode to run")
)

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	log.Printf("mode type: %T, address: %p, value: %d", mode, mode, *mode)

	switch *mode {
	case MODE_HELLO:
		demo.MakeHello()
	case MODE_REQUEST:
		demo.MakeRequest()
	case MODE_LOG:
		demo.MakeLog()
	case MODE_REST:
		demo.MakeRest()
	}
}
