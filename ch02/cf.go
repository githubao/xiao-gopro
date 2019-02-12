// 导入其他包的行为
// author: baoqiang
// time: 2019/2/12 下午8:49
package ch02

import (
	"os"
	"strconv"
	"fmt"
	"github.com/githubao/xiao-gopro/ch02/tempconv"
)

func CF() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))

	}
}
