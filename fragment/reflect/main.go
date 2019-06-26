package main

import (
	"fmt"
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStructure(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

// init：包级别函数，会在包被第一次加载的时候运行
// 作用：
// - 初始化那些不能被初始化表达式完成初始化的变量
// - 检查或者修复程序的状态
// - 注册
// - 仅执行一次的计算
// - 更多其它场合
//
// 包的初始化
// 要想使用导入的包首先需要初始化它，这是由 golang 的运行系统完成的，主要包括(顺序很重要)：
// - 初始化导入的包（递归的定义）
// - 在包级别为声明的变量计算并分配初始值
// - 执行包内的 init 函数
// *重要*不管包被导入多少次，都只会被初始化一次。
// ?貌似?同一包内，不同文件中的 init 会按文件名顺序执行
// 其他
// 在同一个包或者文件当中可以定义很多的 init 函数
// 仅仅为了使用包的副作用（包的初始化）而导入包，比如 import _ "image/png"
func init() {
	// 注册 Bar 到类型注册表
	registerType((*Bar)(nil))
}

type Bar struct {
	Name string
	Sex  int
}

func main() {

	//registerType((*Bar)(nil))
	// 结构体名称
	structName := "Bar"

	// 利用反射生成一个 Bar
	s, ok := newStructure(structName)
	if !ok {
		return
	}

	fmt.Println(s, reflect.TypeOf(s))

	t, ok := s.(Bar)
	if !ok {
		return
	}
	t.Name = "i am bar"

	fmt.Println(t, reflect.TypeOf(t))
}
