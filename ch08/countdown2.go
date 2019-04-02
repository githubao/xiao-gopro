// 计数服务
// author: baoqiang
// time: 2019-04-01 21:50
package ch08

import (
	"fmt"
	"os"
	"time"
)

func Countdown2() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
	//pass
	case <-abort:
		fmt.Println("Lauch aborted!")
		return
	}

	launch()
}
