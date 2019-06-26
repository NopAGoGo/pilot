package main

import (
	"encoding/json"
	"fmt"
)

var stringMap = map[string]string{}

func main() {
	stringMap["file"] = "info"
	stringMap["hello"] = "world"
	fmt.Printf("%v\n", stringMap)

	v := stringMap["world"]
	fmt.Printf("%v\n", v)
	bytes, _ := json.Marshal(stringMap)
	fmt.Println(string(bytes))
}
