// 嵌套的结构体
// author: baoqiang
// time: 2019/3/4 下午7:26
package ch04

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point  Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func RunEmbed() {
	var w Wheel
	//!+
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)

	w.Circle.Point.X = 42

	fmt.Printf("%#v\n", w)
}
