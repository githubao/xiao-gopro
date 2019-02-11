// 读取文件输入
// author: baoqiang
// time: 2019/2/10 下午9:34
package ch01

import (
	"os"
	"bufio"
	"fmt"
)

func Dup02(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin,counts)
	}else {
		for _,arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2: %v\n", err)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}