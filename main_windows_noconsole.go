//go:build windows && noconsole

// 若用vscode编辑该文件，请将 settings.json 的 gopls 的 build.buildFlags 取消注释

package main

import (
	"os/exec"
)

func main() {
	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		defer panic(r)
	// 		errMsg := fmt.Sprint(r)
	// 		walk.MsgBox(nil, "jtrclient-launcher", errMsg, walk.MsgBoxIconError)
	// 	}
	// }()

	myInit()

	cmd := exec.Command(`javaw.exe`, getArgs()...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}
