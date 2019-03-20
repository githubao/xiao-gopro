// go程 并发服务
// author: baoqiang
// time: 2019/3/20 下午7:10
package ch08

import (
	"net"
	"log"
)

func RunServer2() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}

}
