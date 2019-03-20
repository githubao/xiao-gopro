// server响应客户端的请求
// author: baoqiang
// time: 2019/3/20 下午7:14
package ch08

import (
	"net"
	"log"
	"github.com/cyfdecyf/bufio"
	"time"
	"fmt"
	"strings"
)

//Hello
//	HELLO
//	Hello
//	hello
func RunServer3() {
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

		go handleConn3(conn)
	}

}
func handleConn3(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
