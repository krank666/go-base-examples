package main

import "fmt"

const DEFALUT int = 999
const TEAT = "STRING"

var a string = "123"

var f int = DEFALUT

func main(){
	C:=TEAT
	fmt.Printf("%d°DEFALUT or %d°f or %s°c\n",DEFALUT, f, a)
	fmt.Println(C)
}