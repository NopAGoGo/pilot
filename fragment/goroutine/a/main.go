package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 这个可以理解为最多有多少 goroutine 同时存在
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		go prt(wg, i)
	}
	wg.Wait()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func prt(wg *sync.WaitGroup, i int) {
	wg.Add(1)
	fmt.Println("******************")
	fmt.Println("goroutine",i,"创建")
	defer func(){
		wg.Done()
		fmt.Println("goroutine",i,"退出")
		fmt.Println("******************")
	}()
}
