// 测试文件
// author: baoqiang
// time: 2019/2/10 下午8:38
package ch01

import "testing"

func TestRun01(t *testing.T) {

	if true {
		//pass
	}

	if false {
		HelloWorld()
		Echo01()
		Echo02()
		Echo03()
		Dup01()
		Dup02()
		Dup03()
		Lissajous()
		Fetch()
		FetchAll()
		Server1()
		Server2()
		Server3()
	}

}
