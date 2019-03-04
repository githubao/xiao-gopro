// map实现set去除重复行
// author: baoqiang
// time: 2019/3/4 下午6:42
package ch04

import (
	"github.com/cyfdecyf/bufio"
	"os"
	"fmt"
)

func Dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

}
