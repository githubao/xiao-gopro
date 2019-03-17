// 更改struct的值
// author: baoqiang
// time: 2019/3/17 下午8:01
package ch06

import "fmt"

type Values map[string][]string

func RunValues() {
	m := Values{"lang": {"en"}}

	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))

	fmt.Println(m["item"])

	m = nil

	fmt.Println(m.Get("item"))
	//m.Add("item", "3")
}

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
