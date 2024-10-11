package main

import "flag"

var argNoparse = flag.Bool("noparse", false, "Exit directly when the child process crashes.")
var argDelay = flag.Int64("delay", 0, "Dealy start (millisecond).")

func parseArgs() {
	flag.Parse()
}
