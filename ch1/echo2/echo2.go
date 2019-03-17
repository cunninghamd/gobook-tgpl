package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// other ways to declare:
	// s := ""
	// var s string
	// var s = ""
	// var s string = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
