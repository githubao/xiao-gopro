// 使用互斥锁避免竞争条件
// author: baoqiang
// time: 2019-04-08 21:27
package ch09

// go test -run="TestRun09" -race -v
func RunMemo02() {
	m := New(HTTPGetBody)
	//Sequential(m)
	Concurrent(m)
}

func (m *Memo) Get2(key string) (interface{}, error) {
	m.mu.Lock()

	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}

	m.mu.Unlock()

	return res.value, res.err
}
