// slice的底层操作
// author: baoqiang
// time: 2019/3/4 下午6:23
package ch04

import "fmt"

func RunNonempty(){
	//s := []string{"1","a","b"}
	s := []string{"1","","b"}
	fmt.Println(nonempty(s))
}



func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
