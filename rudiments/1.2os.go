package main

import (
	"fmt"
	"os"
)


func main(){
	TestOs()
}

func TestOs(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args)
	fmt.Println(s)
}
