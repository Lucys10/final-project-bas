package internal

import "sync"

type MapCache struct {
	Mpc map[string]struct{}
	mu  sync.Mutex
}

func (m *MapCache) add(fileName string) {
	m.mu.Lock()
	m.Mpc[fileName] = struct{}{}
	m.mu.Unlock()
}

func (m *MapCache) delete(fileName string) {
	m.mu.Lock()
	delete(m.Mpc, fileName)
	m.mu.Unlock()
}

func (m *MapCache) chekAddFile(fileName string) bool {
	m.mu.Lock()
	_, ok := m.Mpc[fileName]
	m.mu.Unlock()
	if ok {
		return ok
	}
	return false
}
