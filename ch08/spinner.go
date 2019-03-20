// 演示goroutine的用法
// author: baoqiang
// time: 2019/3/20 下午5:38
package ch08

import (
	"time"
	"fmt"
)

func RunSpinner() {
	go spinner(1000 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(dealy time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(dealy)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
