// defer记录函数时间
// author: baoqiang
// time: 2019/3/13 上午11:16
package ch05

import (
	"log"
	"time"
)

func RunTimeDefer() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()

	log.Printf("enter %s", msg)

	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
