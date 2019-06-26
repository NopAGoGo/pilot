package main

import "fmt"

// 接口
type Person interface {
	SelfIntroduce()
}

// 结构体
type American struct {
	name string

	profile struct {
		age    int
		gender string
	}
}

func main() {
	sam := American{name: "sam"}
	sam.profile.age = 20
	sam.profile.gender = "male"

	PrintfHere("%p", &sam)  // 指针，十六进制方式，带 0x 打印，比如：0xc0000401c0
	PrintfHere("%#p", &sam) // 指针，十六进制方式，不带 0x 打印，比如：c0000401c0
	PrintfHere("%v", sam)   // 正常打印，比如：{sam {20 male}}
	PrintfHere("%+v", sam)  // 带字段名称打印，比如：{sam {12345 67890}}
	PrintfHere("%#v", sam)  // 以 Go 的语法格式打印
	PrintfHere("%T", sam)   // 打印变量的类型或函数签名

	fmt.Printf("\n")

	PrintfHere("%b", 30) // 整型以二进制方式打印
	PrintfHere("%o", 30) // 整型以八进制方式打印
	PrintfHere("%d", 30) // 整型以十进制方式打印
	PrintfHere("%x", 30) // 整型以十六进制方式打印

	fmt.Printf("\n")

	PrintfHere("|%5d|", 123456789) // 最大长度5的整型，超出不截断，不超出则用空格补齐5位并右对齐
	PrintfHere("|%5d|", 1)         // 最大长度5的整型，超出不截断，不超出则用空格补齐5位并右对齐
	PrintfHere("|%-5d|", 1)        // 最大长度5的整型，超出不截断，不超出则用空格补齐5位并左对齐
	PrintfHere("|%05d|", 1)        // 最大长度5的整型，超出不截断，不超出则用0补齐5位并右对齐

	fmt.Printf("\n")

	PrintfHere("%d", 30)            // 整型
	PrintfHere("%+d", 30)           // 带符号的整型
	PrintfHere("%f", 19.01234551)   // 浮点数，默认6位小数，四舍五入
	PrintfHere("%.3f", 19.01234551) // 浮点数，3位小数，四舍五入

	PrintfHere("|%6.3f|", 19.01234551) // 最小6个宽度（包含小数点)，3位小数，四舍五入，右对齐，超过6位时，不会截断
	PrintfHere("|%6.3f|", 9.01) // 最小6个宽度（包含小数点)，3位小数，四舍五入，右对齐，超过6位时，不会截断
	PrintfHere("|%6.2f|", 9.01) // 最小6个宽度（包含小数点)，2位小数，四舍五入，右对齐，超过6位时，不会截断
	PrintfHere("|%6.1f|", 9.01) // 最小6个宽度（包含小数点)，1位小数，四舍五入，右对齐，超过6位时，不会截断
	PrintfHere("|%-6.1f|", 9.01) // 最小6个宽度（包含小数点)，1位小数，四舍五入，左对齐，超过6位时，不会截断

	fmt.Printf("\n")

	PrintfHere("%e", 19.01234551)   // 科学计数法，参数需要浮点数，默认6位小数，四舍五入，e小写
	PrintfHere("%.3e", 19.01234551) // 科学计数法，参数需要浮点数，3位小数，四舍五入，e小写
	PrintfHere("%E", 19.01234551)   // 科学计数法，参数需要浮点数，默认6位小数，四舍五入，e大写
	PrintfHere("%.3E", 19.01234551) // 科学计数法，参数需要浮点数，3位小数，四舍五入，e大写
	PrintfHere("%.6g", 19.01234551) // 浮点数，从左往右取6位，总长6位，四舍五入，.0会被省略

	fmt.Printf("\n")

	PrintfHere("%x", 30)  // 整型以十六进制、字母小写方式显示，不带空格
	PrintfHere("% x", 30) // 整型以十六进制、字母小写方式显示，带空格
	PrintfHere("%X", 30)  // 整型以十六进制、字母大写方式显示，不带空格
	PrintfHere("% X", 30) // 整型以十六进制、字母大写方式显示，带空格

	fmt.Printf("\n")

	PrintfHere("%x", "hex this")  // 如果参数是字符串，则打印字符串的每一个字符的 ASCII 码
	PrintfHere("% x", "hex this") // 如果参数是字符串，则打印字符串的每一个字符的 ASCII 码
	PrintfHere("%X", "hex this")  // 如果参数是字符串，则打印字符串的每一个字符的 ASCII 码
	PrintfHere("% X", "hex this") // 如果参数是字符串，则打印字符串的每一个字符的 ASCII 码

	fmt.Printf("\n")

	PrintfHere("%U", 'A')                  // 打印 Unicode 字符
	PrintfHere("%s", "你好")                 // 字符串
	PrintfHere("%q", "你好")                 // 给字符串打上双引号
	PrintfHere("%#q", "你好")                // 给字符串打上反引号，如果字符串内有反引号，就用双引号代替
	PrintfHere("%#q", "`你好`")              // 给字符串打上反引号，如果字符串内有反引号，就用双引号代替，相当于 %q
	PrintfHere("|%.5s|", "你好我是梁朝伟")        // 最大宽度为5，超出截断
	PrintfHere("|%5s|", "你好")              // 最小宽度为5，右对齐
	PrintfHere("|%-5s|", "你好")             // 最小宽度为5，左对齐
	PrintfHere("|%05s|", "你好")             // 如果宽度小于5，就会在字符串前面补零
	PrintfHere("|%5.7s|", "你好")            // 最小宽度为5，最大宽度为7，右对齐，超出截断
	PrintfHere("|%-5.7s|", "你好")           // 最小宽度为5，最大宽度为7，左对齐，超出截断
	PrintfHere("|%5.7s|", "你好我是梁朝伟，德华是你吗") // 最小宽度为5，最大宽度为7，右对齐，超出截断
	PrintfHere("|%5.3s|", "你好我是")          // 显示最小宽度为5，右对齐，如果字符宽度大于3，则截断
	PrintfHere("|%-5.3s|", "你好我是")         // 显示最小宽度为5，左对齐，如果字符宽度大于3，则截断

	fmt.Printf("\n")

	PrintfHere("%t", true) // 打印 true 或 false
}

func PrintfHere(format string, a interface{}) {
	fmt.Printf("%-8s%s"+format+"\n", format, ": ", a)
}
