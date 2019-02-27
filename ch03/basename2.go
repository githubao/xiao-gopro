// basename 使用index函数
// author: baoqiang
// time: 2019/2/27 下午7:31
package ch03

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

// a/b/c.d.go
// c.d
func RunBasename2() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename2(input.Text()))
	}
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}
