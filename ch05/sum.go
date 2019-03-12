// 可变参数
// author: baoqiang
// time: 2019/3/12 下午9:27
package ch05

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func RunSum() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))

	vals := []int{4, 5, 6}
	fmt.Println(sum(vals...))
}
