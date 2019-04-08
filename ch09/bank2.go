// 互斥操作获得锁
// author: baoqiang
// time: 2019-04-08 20:47
package ch09

import (
	"fmt"
	"sync"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func RunBank02() {
	var n sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit2(amount)
			n.Done()
		}(i)
	}

	n.Wait()

	if got, want := Balance2(), (1001*1000)/2; got != want {
		fmt.Printf("balance = %d, want %d", got, want)
	}

}

func Deposit2(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
