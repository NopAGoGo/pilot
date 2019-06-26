package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// done
func main() {

	var n int64 = 511 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]

	fmt.Println(Int64ToBytes(n))

	var b []byte = []byte{6: 1, 7: 255}

	fmt.Println(BytesToInt64(b))
}

func Int64ToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	_ = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
