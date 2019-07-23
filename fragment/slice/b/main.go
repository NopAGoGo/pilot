package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ptr unsafe.Pointer // Array pointer
	len int            // slice length
	cap int            // slice capacity
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := append(s1, 10)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	s2[0] = 0
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	// 把slice转换成自定义的 Slice struct
	slice1 := (*Slice)(unsafe.Pointer(&s1))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice1.ptr, slice1.len, slice1.cap)
	slice2 := (*Slice)(unsafe.Pointer(&s2))
	fmt.Printf("ptr:%v len:%v cap:%v \n", slice2.ptr, slice2.len, slice2.cap)
}
