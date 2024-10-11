package main

import "os"

const syslibNewPath = "syslib_new"
const binPath = "bin"
const binNewPath = "bin_new"

func completeUpgrade() {
	println("[complete upgrade]")

	c := func(newPath, oldPath string) {
		println(newPath, "=>", oldPath)

		newInfo, err := os.Stat(newPath)
		if err != nil || !newInfo.IsDir() {
			return
		}
		err = os.RemoveAll(oldPath)
		if err != nil {
			panic(err)
		}
		err = os.Rename(newPath, oldPath)
		if err != nil {
			panic(err)
		}
	}

	c(syslibNewPath, syslibPath)
	c(binNewPath, binPath)

	println()
}
