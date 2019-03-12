// 乘方
// author: baoqiang
// time: 2019/3/12 下午8:56
package ch05

import "fmt"

func RunSquares() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
