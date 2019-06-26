package main

import (
	"fmt"
	"math/rand"
	"time"
)

// done
func main() {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		fmt.Println(rand.Int())
	}
}
