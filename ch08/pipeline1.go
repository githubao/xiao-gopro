// 管道数据流
// author: baoqiang
// time: 2019/3/20 下午7:50
package ch08

import "fmt"

func RunPipeline1() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		//for x := 0; x <= 10; x++ {
		for x := 0;; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}

}
