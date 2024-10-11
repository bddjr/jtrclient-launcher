package main

import (
	"os"
	"path/filepath"
	"strings"
)

func setEnv() {
	const envName = "PATH"
	println("[" + envName + "]")
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ls := []string{
		filepath.Join(dir, "bin"),
		filepath.Join(dir, "VLC"),
		filepath.Join(dir, "jdk", "bin"),
		filepath.Join(dir, "ant", "bin"),
		filepath.Join(dir, "ffmpeg", "bin"),
		filepath.Join(dir, "Python"),
		os.Getenv(envName),
	}
	err = os.Setenv(envName, strings.Join(ls, delimiter))
	if err != nil {
		panic(err)
	}
	print(os.Getenv(envName), "\n\n")
}
