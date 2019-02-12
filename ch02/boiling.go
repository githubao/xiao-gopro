// 常量声明
// author: baoqiang
// time: 2019/2/12 下午8:27
package ch02

import "fmt"

const boilingF = 212.0

func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
