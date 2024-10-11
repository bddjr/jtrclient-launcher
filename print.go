//go:build !noconsole

package main

import (
	"fmt"
)

func println(a ...any) {
	fmt.Println(a...)
}

func print(a ...any) {
	fmt.Print(a...)
}
