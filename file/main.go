package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)
	l.PushBack(7)
	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

