/* currying 柯里化
我感觉是自己折腾自己
用到的知识有匿名函数、闭包、高阶函数

高阶函数：可以接收函数作为参数的函数
匿名函数：没有函数名（?私以为? lambda 表达式：一般有箭头，可以实现匿名函数，go 里应该没这玩意）
闭包：函数返回函数，相同作用域内的变量有效

柯里化（Currying）是把接受多个参数的函数变换成接受一个单一参数（最初函数的第一个参数）的函数，并且返回接受余下的参数且返回结果的新函数的技术。
下面是几个例子，写完我自己都晕了
讲道理，以我现在浅薄的知识，感觉没啥应用场景
想到的作用应该是复用代码。。。
*/
package main

import "fmt"

func main() {

	// 实现加法
	s1 := add(3 )(4)
	fmt.Println(s1)

	// 这里是实现了乘法，但是参数跟参数处理方法可以随意传，第一次传了两个参数，感觉不够柯里于是看下面。。。
	s2 := mul(3, func(x, y int) int { return x * y })(4)
	fmt.Println(s2)

	// 这里还是实现了乘法，可能比较柯里了，但是已经不是给人看的了，猜测都是分开写完，然后有助复用代码，可能有点 endpoint 的意思
	s3 := multiple(3)(4)(func(x, y int) int { return x * y })
	fmt.Println(s3)
}


func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func mul(x int, f func(int, int) int) func(y int) int {
	return func(y int) int {
		return f(x, y)
	}
}

func multiple(x int) func(int) func(func(int,int) int) int {
	return func(y int) func(func(int,int) int) int {
		return func(f func(int,int) int) int{
			return f(x,y)
		}
	}
}