// select 多路复用
// author: baoqiang
// time: 2019-04-01 21:54
package ch08

import (
	"fmt"
	"os"
	"time"
)

func Countdown3() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		//pass
		case <-abort:
			fmt.Println("Lauch aborted!")
			return
		}
	}

	launch()
}
