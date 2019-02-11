// 直接使用strings的join方法
// author: baoqiang
// time: 2019/2/10 下午8:59
package ch01

import (
	"fmt"
	"strings"
	"os"
)

func Echo03() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
