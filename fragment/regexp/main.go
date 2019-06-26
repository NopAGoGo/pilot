package main

import (
	"fmt"
	"regexp"
)

// todo 这个可能完整的要写很多。。。
func main() {
	// 判断输入的批号是否有效
	//str :="H906188-1-1"
	//reg := regexp.MustCompile(`[A-Za-z]?[0-9]{5,7}(-[0-9])*`)

	// 挑出第一段中文
	str :="hello你好haha开心"
	reg:=regexp.MustCompile(`[\p{Han}]+`)
	fmt.Println(reg.MatchString(str))
	result := reg.FindString(str)

	fmt.Println(str)
	fmt.Println(result)
}
