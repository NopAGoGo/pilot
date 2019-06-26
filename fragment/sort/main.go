package main

import (
	"fmt"
)

func main(){
	Bubble()
}

func Bubble(){
	sq:=[]int{5,9,3,1,2,8,4,7,6}
	for i:=0;i<9;i++{
		for j:=0;j<9-i-1;j++{
			if sq[j]>sq[j+1]{
				sq[j],sq[j+1]=sq[j+1],sq[j]
			}
		}
	}
	fmt.Println(sq)
}