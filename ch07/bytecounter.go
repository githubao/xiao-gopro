// 实现一个接口
// author: baoqiang
// time: 2019/3/18 下午2:22
package ch07

import "fmt"

type ByteCounter int

func RunByteCounter() {
	var c ByteCounter

	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
