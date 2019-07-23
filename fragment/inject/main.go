package main

import (
	"fmt"
)

// 设计模式：访问者模式
func main() {
	e:=new(Element)
	e.Accept(new(ProductionVisitor))
}

type IVisitor interface{
	Visit()
}

type ProductionVisitor struct{
}

func (v ProductionVisitor) Visit(){
	fmt.Println("这是生产环境")
}

type TestingVisitor struct {
}

func (v TestingVisitor) Visit(){
	fmt.Println("这是测试环境")
}

type Element struct {
}

func (el Element) Accept(visitor IVisitor){
	visitor.Visit()
}

