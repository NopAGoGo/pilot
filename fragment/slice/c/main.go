package main

import (
	"fmt"
	"unsafe"
)



func main() {
	// 定义切片三种方法
	// 方法1
	arr := [...]int{1, 2, 3, 4, 5, 6} // arr为数组,展开运算符...此处作为略写数组的长度
	slice := arr[:]                   // slice引用arr

	// 方法2
	slice2 := make([]int, 6, 12) // make方式创建的切片对应的数组是由make底层维护,对外不可见,也就是只能通过slice访问各个元素
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	slice2[3] = 4
	slice2[4] = 5
	slice2[5] = 6

	// 方法3
	slice3 := []int{1, 2, 3, 4, 5, 6}
	slice3[0] = 1


	// 切片的三种状态
	// 零切片,slice2就是零切片
	sliceZero:=make([]int,10)
	sliceZero[0]=0

	// 空切片
	// nil切片
	sliceNil:=make([]*int,10)
	sliceNil[0]=nil

	slicePointer := &slice
	sliceCopy := slice
	//sliceCopy[0] = 99
	//(*slicePointer)[0] = 98
	//slice = slice[:2] //

	//var slicesct = slice
	fmt.Println(slice)


	sct := *(*[3]int)(unsafe.Pointer(&slice))
	fmt.Printf("test %T \n", sct)
	fmt.Println(sct)


	fmt.Printf("=================================\n")
	fmt.Printf("arr := [...]int{1, 2, 3, 4, 5, 6} \n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("%%v of  arr: %v \n", arr)
	fmt.Printf("%%v of &arr: %v \n", &arr)
	fmt.Printf("%%p of &arr: %p \n", &arr)
	fmt.Printf("%%T of  arr: %T \n", arr)
	fmt.Printf("%%T of &arr: %T \n", &arr)
	fmt.Printf("=================================\n")

	fmt.Printf("\n=================================\n")
	fmt.Printf("slice := []int{1, 2, 3, 4, 5, 6} \n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("%%v of  slice: %v \n", slice)
	fmt.Printf("%%v of &slice: %v \n", &slice)
	fmt.Printf("%%p of  slice: %p \n", slice)
	fmt.Printf("%%p of &slice: %p \n", &slice)
	fmt.Printf("%%T of  slice: %T \n", slice)
	fmt.Printf("%%T of &slice: %T \n", &slice)
	fmt.Printf("=================================\n")

	fmt.Printf("\n=================================\n")
	fmt.Printf("slicePointer := &slice \n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("%%v of *slicePointer: %v \n", *slicePointer)
	fmt.Printf("%%v of  slicePointer: %v \n", slicePointer)
	fmt.Printf("%%v of &slicePointer: %v \n", &slicePointer)
	fmt.Printf("%%p of  slicePointer: %p \n", slicePointer)
	fmt.Printf("%%p of &slicePointer: %p \n", &slicePointer)
	fmt.Printf("%%T of  slicePointer: %T \n", slicePointer)
	fmt.Printf("%%T of &slicePointer: %T \n", &slicePointer)
	fmt.Printf("=================================\n")

	fmt.Printf("\n=================================\n")
	fmt.Printf("sliceCopy := i\n")
	fmt.Printf("---------------------------------\n")
	fmt.Printf("%%v of  sliceCopy: %v \n", sliceCopy)
	fmt.Printf("%%v of &sliceCopy: %v \n", &sliceCopy)
	fmt.Printf("%%p of  sliceCopy: %p \n", sliceCopy)
	fmt.Printf("%%p of &sliceCopy: %p \n", &sliceCopy)
	fmt.Printf("%%T of  sliceCopy: %T \n", sliceCopy)
	fmt.Printf("%%T of &sliceCopy: %T \n", &sliceCopy)
	fmt.Printf("=================================\n")
}
