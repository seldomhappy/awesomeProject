package keyval

import "testing"

func setN(store *MemStore, N int, done chan<- struct{}) {
	for i := 0; i < N; i++ {
		store.Set(i, -i)
	}
	done <- struct{}{}
}

func BenchmarkMemStore(b *testing.B) {
	store := NewMemStore()
	done := make(chan struct{})
	b.ResetTimer()
	go setN(store, b.N, done)
	go setN(store, b.N, done)
	<-done
	<-done
}
