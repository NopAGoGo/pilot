package main

import "fmt"

type person struct {
	name     string
	oddYears []int
}

func main() {
	sum1 := foo(func(int) {})

	fmt.Println("sum1:", sum1)

	p := &person{
		name: "tom",
	}

	p2, sum2 := bar(p, func(i int) { p.oddYears = append(p.oddYears, i) })
	fmt.Printf("sum2: %d | oddYears: %#v\n", sum2, p2.oddYears)
}

func foo(f func(i int)) int {
	sum := 0
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			sum += i
			f(i)
		}
	}
	sum *= sum
	return sum
}

func bar(p *person, f func(i int)) (*person, int) {
	//p := &person{
	//  name: "tom",
	//}
	//f := func(i int) {
	//  p.oddYears = append(p.oddYears, i)
	//}
	sum := foo(f)
	return p, sum
}
