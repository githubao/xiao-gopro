// 两步锁
// author: baoqiang
// time: 2019-04-08 21:32
package ch09

// go test -run="TestRun09" -race -v
func RunMemo03() {
	m := New(HTTPGetBody)
	//Sequential(m)
	Concurrent(m)
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	m.mu.Unlock()

	if !ok {
		res.value, res.err = m.f(key)

		m.mu.Lock()
		m.cache[key] = res
		m.mu.Unlock()
	}

	return res.value, res.err
}
