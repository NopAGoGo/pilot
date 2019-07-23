package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type Person struct {
	Name string
}

func (p Person) Hello() {
	for {
		fmt.Println(p.Name)
		time.Sleep(time.Second * 1)
	}
}

func (p *Person) Rename() {
	time.Sleep(time.Second * 5)
	p.Name = "sunday"
	fmt.Println("name changed")
}

func main() {
	runtime.GOMAXPROCS(2)
	p := Person{"tom"}
	go p.Hello()
	//time.Sleep(time.Second*20)
	go p.Rename()
	fmt.Println("**********")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		fmt.Println("hmm")
		sig := <-ch
		switch sig {
		case syscall.SIGTERM:
			// stop
			fmt.Println("stop")
			log.Printf("stop")
			signal.Stop(ch)
			fmt.Println("graceful shutdown")
			log.Printf("graceful shutdown")
			return
		}
	}
}
