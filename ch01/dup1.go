// stdin读取输入
// author: baoqiang
// time: 2019/2/10 下午9:19
package ch01

import (
	"github.com/cyfdecyf/bufio"
	"os"
	"fmt"
)

func Dup01() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
