// 反转数组
// author: baoqiang
// time: 2019/2/27 下午8:20
package ch04

import "fmt"

func RunReverse() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverse(a[:])
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
