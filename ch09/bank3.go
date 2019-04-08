// sync锁
// author: baoqiang
// time: 2019-04-08 20:53
package ch09

import (
	"fmt"
	"sync"
)

var (
	mu       sync.Mutex
	balance3 int
)

func RunBank03() {
	var n sync.WaitGroup

	for i := 1; i <= 500; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit3(amount)
			n.Done()
		}(i)
	}

	n.Wait()

	if got, want := Balance3(), (501*500)/2; got != want {
		fmt.Printf("balance = %d, want %d", got, want)
	}

}

func Deposit3(amount int) {
	mu.Lock()
	balance3 += amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	b := balance3
	mu.Unlock()
	return b
}
