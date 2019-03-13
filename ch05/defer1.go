// defer包裹调用
// author: baoqiang
// time: 2019/3/13 上午11:28
package ch05

import "fmt"

// f(3)
// f(2)
// f(1)
// defer 1
// defer 2
// defer 3
func RunDefer() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
