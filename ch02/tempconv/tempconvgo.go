// 调用里面的函数
// author: baoqiang
// time: 2019/2/12 下午10:02
package tempconv

import "fmt"

func TempconvGo() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	fmt.Printf("%g\n", CtoF(BoilingC)-CtoF(FreezingC))

	c := FtoC(212.0)
	fmt.Println(c)
	fmt.Println(c.String())
	fmt.Printf("%v\n",c)
	fmt.Printf("%s\n",c)
	fmt.Printf("%g\n",c)
	fmt.Println(float64(c))
}
