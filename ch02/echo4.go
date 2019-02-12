// 使用flag包传递参数
// author: baoqiang
// time: 2019/2/12 下午8:35
package ch02

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit training newline")
var sep = flag.String("s", " ", "separator")

// ./main -s / -n a b cd
// a/b/cd%
func Echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
