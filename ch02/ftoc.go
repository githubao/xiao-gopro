// 声明一个函数
// author: baoqiang
// time: 2019/2/12 下午8:31
package ch02

import "fmt"

func Ftoc() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))

}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
