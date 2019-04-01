// 网络连接
// author: baoqiang
// time: 2019/3/20 下午5:45
package ch08

import (
	"io"
	"log"
	"net"
	"time"
)

func RunServer1() {
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

		handleConn(conn)
	}

}

//Mon Jan 2 03:04:05PM 2006 UTC­0700
//time.RFC3339
func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
