// 有颜色的点
// author: baoqiang
// time: 2019/3/17 下午8:07
package ch06

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func RunColoredPoint() {
	RunColoredPoint1()
	RunColoredPoint2()
	RunColoredPoint3()
}

func RunColoredPoint3() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{Point{1, 2}, red}
	q := ColoredPoint{Point{4, 6}, blue}

	fmt.Println(p.Distance(q.Point))

	p.ScaleBy(2)
	q.ScaleBy(2)

	fmt.Println(p.Distance(q.Point))

	fmt.Println()
}

func RunColoredPoint2() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint2{&Point{1, 2}, red}
	q := ColoredPoint2{&Point{4, 6}, blue}

	fmt.Println(p.Distance(*q.Point))

	q.Point = p.Point

	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)

	fmt.Println()
}

func RunColoredPoint1() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance

	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	//scale := Point.ScaleBy // err
	//scale(p, 2)

	fmt.Println(p)
	fmt.Printf("%T\n", scale)

	fmt.Println()
}
