//go:build unix

package main

import (
	"os"
	"syscall"
)

func main() {
	myInit()
	err := syscall.Exec(lookPath(`java`), getArgs(), os.Environ())
	panic(err)
}
