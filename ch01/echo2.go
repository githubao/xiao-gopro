// echo2
// author: baoqiang
// time: 2019/2/10 下午8:56
package ch01

import (
	"os"
	"fmt"
)

func Echo02() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
