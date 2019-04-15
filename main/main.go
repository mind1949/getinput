package main

import (
	"fmt"
	"github.com/mind1949/getinput"
)

func main() {
	v1 := getinput.Call("string:", "string")
	fmt.Printf("v1=> type: %T, value:  %v\n",v1, v1)

	v2 := getinput.Call("int:", "int")
	fmt.Printf("v2=> type: %T, value:  %v\n",v2, v2)

	v3 := getinput.Call("float:", "float64")
	fmt.Printf("v3=> type:  %T, value: %v\n",v3, v3)
}
