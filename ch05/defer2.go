// recover
// author: baoqiang
// time: 2019/3/13 上午11:31
package ch05

import (
	"os"
	"runtime"
)

func RunRecover() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte

	n := runtime.Stack(buf[:], false)

	os.Stdout.Write(buf[:n])
}
