// 包里面使用常量具体的函数实现
// author: baoqiang
// time: 2019/2/12 下午8:46
package tempconv

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
