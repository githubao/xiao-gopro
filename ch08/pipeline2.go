// 关闭一个chan
// author: baoqiang
// time: 2019/3/20 下午7:54
package ch08

import "fmt"

func RunPipeline2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}

		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}

}
