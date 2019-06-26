package main

import (
	"fmt"
)

// done
// 用于测试是否每次循环都会计算len(intSlice)
// 答案是会
func main() {

	// 有8个元素
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7}

	// 第6个元素被删除，实际循环了7次
	for i := 0; i < len(intSlice); i++ {
		if intSlice[i] == 5 {
			intSlice = append(intSlice[:i], intSlice[i+1:]...)
		}
		fmt.Println(i)
	}
}
