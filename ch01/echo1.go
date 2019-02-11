// for循环
// author: baoqiang
// time: 2019/2/10 下午8:40
package ch01

import (
	"os"
	"fmt"
)

func Echo01() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
