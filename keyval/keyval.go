package keyval

import (
	"sync"
)

type MemStore struct {
	data  map[interface{}]interface{}
	mutex sync.RWMutex
}

func NewMemStore() *MemStore {
	m := &MemStore{
		data: make(map[interface{}]interface{}),
		// mutex does not need initializing
	}
	return m
}

func (m *MemStore) Set(key interface{}, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if value != nil {
		m.data[key] = value
	} else {
		delete(m.data, key)
	}
}
