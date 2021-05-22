package main

import (
	"fmt"
	"github.com/facebookgo/inject"
)

type A struct {
	Name string
}


type Store struct {
	DBO *A
	BBO *A
}

var injectGraph inject.Graph

func main() {
	a := A{
		Name: "hello",
	}
	a2 := A{
		Name: "hello2",
	}

	c := Store{}
		err := injectGraph.Provide(//对象提供者
		&inject.Object{Value: &c},//这个也需要
		&inject.Object{Value: &a},
		&inject.Object{Value: &a2},
	)
	if err != nil{}
	err = injectGraph.Populate()//填充对象到使用了inject标签的结构体字段中
	if err != nil{}
	fmt.Println(c.BBO.Name)
	fmt.Println(c.DBO.Name)
}
// 通过injectGraph.Objects()可以获取所有设置了Name的待填充对象
func GetObject(name string) interface{} {
	for _, o := range injectGraph.Objects() {
		if o.Name == name {
			return o.Value
		}
	}
	return nil
}