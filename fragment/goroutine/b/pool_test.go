package a

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func GoPool() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	data := make(chan int, 100)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _ = range data {
				//fmt.Println(d)
			}
		}()
	}

	for i := 0; i < 10000; i++ {
		data <- i
	}
	close(data)
	wg.Wait()
}

func NoPool() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkGoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoPool()
	}
}

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPool()
	}
}
