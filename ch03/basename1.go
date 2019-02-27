// 获取文件的名字
// author: baoqiang
// time: 2019/2/27 下午6:48
package ch03

import (
	"github.com/cyfdecyf/bufio"
	"os"
	"fmt"
)

func RunBasename1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
