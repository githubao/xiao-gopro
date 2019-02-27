// 打印int数组
// author: baoqiang
// time: 2019/2/27 下午7:45
package ch03

import (
	"bytes"
	"fmt"
)

func RunInts2String() {
	ints := []int{1, 3, 2}

	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range ints {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	fmt.Println(buf.String())
}
