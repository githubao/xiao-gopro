// 并发条件
// author: baoqiang
// time: 2019-04-08 20:38
package ch09

import "fmt"

// 避免并发的三个条件：静态加载，使用channel，锁

var deposits = make(chan int)
var balances = make(chan int)

func RunBank01() {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		fmt.Printf("balance = %d, want %d", got, want)
	}

}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			//pass
		}
	}
}

func init() {
	go teller()
}
