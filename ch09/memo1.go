// 缓存结果函数
// author: baoqiang
// time: 2019-04-08 21:07
package ch09

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// go test -run="TestRun09" -race -v
func RunMemo01() {
	m := New(HTTPGetBody)
	Sequential(m)
	//Concurrent(m)
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func (m *Memo) Get1(key string) (interface{}, error) {
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}
