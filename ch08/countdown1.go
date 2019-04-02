// 统计个数
// author: baoqiang
// time: 2019-04-01 21:45
package ch08

import (
	"fmt"
	"time"
)

func Countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
