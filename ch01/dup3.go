// ioutil直接读取数据
// author: baoqiang
// time: 2019/2/10 下午9:47
package ch01

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func Dup03() {
	counts := make(map[string]int)

		for _,filename := range os.Args[1:]{
			data, err := ioutil.ReadFile(filename)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup3: %v\n", err)
				continue
			}
			for _,line := range strings.Split(string(data),"\n"){
				counts[line]++
			}
		}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
