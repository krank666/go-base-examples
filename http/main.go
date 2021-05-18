package main

import "fmt"

func main() {
	a := make([]int,0,0)
	fmt.Printf("%p,%p\n",a,&a)
	a = append(a,1)
	fmt.Printf("%p,%p",a,&a)
}
