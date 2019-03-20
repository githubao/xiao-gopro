// 实现flag value接口
// author: baoqiang
// time: 2019/3/18 下午3:54
package ch07

import (
	"flag"
	"fmt"
	"time"
)

func RunSleep() {
	flag.Parse()

	fmt.Printf("sleeping for %v...", *period)
	time.Sleep(*period)

	fmt.Println()
}

var period = flag.Duration("period", 1*time.Second, "sleep period")
