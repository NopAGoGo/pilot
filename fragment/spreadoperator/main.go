package main

import (
	"fmt"
	"strings"
)

func join(params ...interface{}) string {
	var paramSlice []string
	for _, param := range params {
		paramSlice = append(paramSlice, param.(string))
	}
	return strings.Join(paramSlice, "_") // Join 方法第2个参数是 string 而不是 rune
}

func main() {
	fmt.Println(join("redis", "100","master"))
}