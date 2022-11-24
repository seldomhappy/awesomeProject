// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	_, err := GetFiles2(context.TODO(), "1", "2")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(time.Since(start))
}

// GetFiles пример функции, которую нужно оптимизировать.
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

var ret sync.Map
var wg sync.WaitGroup

func doMath(a, b int) int {
	if a < b {
		return a + b
	}
	return a - b
}

func GetFiles2(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	l := len(names)
	if l == 0 {
		return nil, nil
	}
	result = make(map[string][]byte, l)

	wg.Add(l)
	errChan := make(chan error, l)

	ch := make(chan struct{})
	for _, name := range names {
		go func(name string) {
			defer wg.Done()

			ch <- struct{}{}
			result[name], err = GetFile(ctx, name)
			<-ch

			data, err := GetFile(ctx, name)
			ret.Store(name, data)
			errChan <- err
		}(name)
	}
	go func() {
		wg.Wait()
		close(ch)
		close(errChan)
	}()
	e := <-errChan

	ret.Range(func(k, v interface{}) bool {
		//v, ok := v.(type)
		result[fmt.Sprint(k)] = []byte(fmt.Sprint(v))
		return true
	})

	fmt.Println(result)

	return result, e
}

// GetFile является примером функции, которая относительно
// недолго выполняется при единичном вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что оптимизировать в этой функции нечего.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
