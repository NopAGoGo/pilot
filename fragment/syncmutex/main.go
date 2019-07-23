package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	//var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
		//go incrementWithMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x is", x)
}

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
