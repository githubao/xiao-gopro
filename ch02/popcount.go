// init函数的使用，统计1的个数
// 使用的是查表法，每隔8位记录的1的个数
// author: baoqiang
// time: 2019/2/12 下午9:42
package ch02

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount() {
	count := popCount(0x1234567890ABCDEF)
	fmt.Println(count)
}

func popCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))] +
			pc[byte(x>>(8*8))])
}

//0000
//0001
//0010
//0011
//0100
//0101
//0110
//0111
//1000
//1001
//1010
//1011
//1100
//1101
//1110
//1111
func PrintHex() {
	for i := 0; i < 16; i++ {
		fmt.Printf("%04b\n", i)
	}
}
