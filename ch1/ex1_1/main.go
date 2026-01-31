package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i, arg := range os.Args[0:] {
		s += fmt.Sprintf("Index: %d argument: %s\n", i, arg)
	}
	fmt.Println(s)
}
