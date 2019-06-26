package main

import "fmt"

// 接口
type Person interface {
	SelfIntroduce()
}

// 结构体
type American struct {
	name string
}

// 方法即函数，方法本身就是一个正常的函数，功能并没有什么变化。

// 方法接收者是值
// 方法与指针重定向，American 实现了 Person 接口
// 带指针参数的函数必须接受一个指针，而以指针为接收者的方法被调用时，接收者既能为值又能为指针。
// 上面那句比较拗口，其实就是方法接收者是结构体值的方法，结构体值与结构体指针都能调用。（可能讲的不对，大概意思就是这样）
// 验证：接收者是值的方法，在执行时，是否会发生拷贝 答案是会
// 调用时，值拷贝一个新的结构体，如有修改，修改的是拷贝的新结构体，接收者不受影响，所以私认为一般这样的方法会产生输出，或有返回值
// 好比 SelfIntroduce(a American)？
func (a American) SelfIntroduce() {
	fmt.Println("hello, my name is", a.name)
	fmt.Printf("自我介绍时，介绍主体的地址:%p，主体name:%s\n",&a,a.name)
}

// 结构体
type Alien struct {
	name string
}

//
// 方法接收者是指针，Alien 也实现了 Person 接口
// todo 待验证，指针才实现了 Person 接口
// 验证：接收者是指针的方法，在执行时，是否会发生拷贝 答案是不会
// 调用时不发生拷贝，如有修改则对指针指向的值做出修改，即指针接收者的方法可以修改指针指向的值
// *重要*方法不能改变指针的指向
// 好比 SelfIntroduce(a *Alien)？
func (a *Alien) SelfIntroduce() {
	fmt.Println("おまえは、もう死んでいる")
	fmt.Printf("自我介绍时，介绍主体的地址:%p，主体name:%s\n",a,a.name)
}

// 判断 American 是否实现了 Person
// todo 三种有什么区别。。。待续
var _ Person = new(American)
var _ Person = &American{}
var _ Person = American{}

func main() {
	var p1,p2 Person
	fmt.Printf("p1初始化后的地址： %p\n",&p1)
	fmt.Printf("p2初始化后的地址： %p\n",&p2)

	tom:= American{"tom"}
	fmt.Printf("tom的地址：%p\n",&tom)
	tom.SelfIntroduce()

	p1 = tom
	fmt.Printf("将tom赋给p1后，p1的地址： %p\n",&p1) // 这里赋值是值拷贝

	jack := &Alien{"jack"}
	fmt.Printf("jack的地址：%p\n",jack)
	jack.SelfIntroduce()

	p2 = jack
	fmt.Printf("将jack赋给p2后，p2的地址： %p\n",&p2) // 这里赋值也是值拷贝，

	p1.SelfIntroduce()

	p2.SelfIntroduce() // 这里厉害了，之前已经把 jack 的指针赋给 p2，p2 调用后，执行方法的主体仍是原来的 jack，原理不明，以后填坑 todo


	if _, ok := p1.(*Alien); ok {
		fmt.Println("OK")
	} else{
		fmt.Println("HMMM...")
	}

	if _, ok := p2.(*Alien); ok {
		fmt.Println("OK")
	}else{
		fmt.Println("HMMM...")
	}
}
