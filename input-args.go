package main

import "flag"

var argDelay = flag.Int64("delay", 0, "Dealy start (millisecond).")

func parseArgs() {
	flag.Parse()
}
