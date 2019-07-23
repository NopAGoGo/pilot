package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	GoPool()
	NoPool()
}

func NoPool() {
	start := time.Now()
	wg := new(sync.WaitGroup)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func GoPool() {
	start := time.Now()
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
	end := time.Now()
	fmt.Println(end.Sub(start))
}
