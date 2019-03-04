// 字符数组的相等性判断
// author: baoqiang
// time: 2019/2/27 下午8:13
package ch04

import (
	"crypto/sha256"
	"fmt"
)

func RunSha256() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
