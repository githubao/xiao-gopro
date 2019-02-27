// 位运算
// author: baoqiang
// time: 2019/2/27 下午7:51
package ch03

// ^ 一元运算符表示按位取反
// ^ 二元运算符表示异或，相同为0不同为1
// &^ 位清空， c = a&^b, b中为1的位置c变成0, 其他位置与a的值相同

type Flags uint

const (
	FlagUp           Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
