package main

import (
	"os/exec"
)

func lookPath(file string) string {
	println("[file]")
	path, err := exec.LookPath(file)
	if err != nil {
		panic(err)
	}
	print(path, "\n\n")
	return path
}
