package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const mainName = "top.jtalkroom.jtrclient.QJTRClient"
const syslibPath = "syslib"
const fileSuffix = ".jar"

const delimiter = string(os.PathListSeparator)

func getArgs() []string {
	println("[args]")
	cp := getCP()
	args := []string{"-classpath", cp, mainName, "0"}
	print(args, "\n\n")
	return args
}

func getCP() string {
	cp := "."
	err := filepath.Walk(syslibPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), fileSuffix) {
			cp += delimiter + path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return cp
}
