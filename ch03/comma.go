// 数字隔1000打印逗号
// author: baoqiang
// time: 2019/2/27 下午7:37
package ch03

import (
	"os"
	"fmt"
)

func RunComma() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}
