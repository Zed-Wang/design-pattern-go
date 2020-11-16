package main

import (
	"design-pattern-go/singleton_1"
	"fmt"
)

func main() {
	single1 := singleton.GetInstance()
	single2 := singleton.GetInstance()
	if single1 == single2 {
		fmt.Println("T")
	}
}
