package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat(".gitignore")
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(fileinfo.Name())    // 获取文件名
	fmt.Println(fileinfo.IsDir())   // 判断是否是目录，返回bool类型
	fmt.Println(fileinfo.ModTime()) // 获取文件修改时间
	fmt.Println(fileinfo.Mode())    // 获取文件权限，-OwnerRWXGroupRWXOthersRWX，读 写 可执行 分别对应的是 r w x 如果没有那一个权限，用 - 代替。第一个 - 代表普通文件，目录的话会是 -d
	fmt.Println(fileinfo.Size())    // 获取文件大小，单位字节，bytes
}
