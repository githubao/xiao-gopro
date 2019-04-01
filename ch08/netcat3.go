// 使用chan保证并发顺序
// author: baoqiang
// time: 2019/3/20 下午7:39
package ch08

import (
	"net"
	"os"
	"log"
)

func RunClient3() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done
}