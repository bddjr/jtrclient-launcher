//go:build !noconsole

package main

import "flag"

var argNoparse = flag.Bool("noparse", false, "Exit directly when the child process crashes.")
