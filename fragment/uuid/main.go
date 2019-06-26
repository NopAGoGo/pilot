package main

import (
	"encoding/hex"
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	uuid, _ := uuid.NewV1()

	fmt.Printf("%v\n", uuid)
	uuidNoHyphen := ridHyphen(uuid)
	fmt.Printf("%v\n", uuidNoHyphen)
}

func ridHyphen(u [16]byte) string {
	buf := make([]byte, 36)
	hex.Encode(buf[:], u[:])
	return string(buf)
}
