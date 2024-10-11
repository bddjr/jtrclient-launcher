//go:build windows && !noconsole

// 若用vscode编辑该文件，请将 settings.json 的 gopls 的 build.buildFlags 注释

package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	myInit()

	// cmd := exec.Command(lookPath(`python.exe`), "-c", `print('hello world');input();raise "???"`)

	cmd := exec.Command(lookPath(`java.exe`), getArgs()...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		if cmd.ProcessState != nil {
			if !*argNoparse {
				print("\n\n Press Ctrl+C to exit: ")
				signalCtrlC := make(chan os.Signal, 1)
				signal.Notify(signalCtrlC,
					// CTRL+C
					syscall.SIGINT,
				)
				<-signalCtrlC
			}
			os.Exit(cmd.ProcessState.ExitCode())
		}
		panic(err)
	}
}
