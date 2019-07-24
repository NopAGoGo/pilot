package main

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/sunday9th/pilot/fragment/protobuf/a/person"
)

func main() {
	msg := &person.Person{
		Name: "sunday",
		Age:  18,
		From: "China",
	}

	// 序列化
	msgDataEncoding, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgEntity := person.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	fmt.Println(msgEntity.GetName())
	fmt.Println(msgEntity.GetAge())
	fmt.Println(msgEntity.GetFrom())

}
