package main

import "fmt"

type MyNum interface {
	type int,float32,float64,string
}

type MyType[T any] struct {

}

func (m MyType[T]) name(t T) (z T) {

}


func myn[T MyNum](t T){
	fmt.Println(t)
}

func main(){
	myn[string]("123")
}