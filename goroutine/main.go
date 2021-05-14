package main

import (
	"fmt"
	"time"
)

func main()  {
	table := make(chan int)
	var Ball int
	for i:=0;i<1000;i++ {
		go player(table,"player")
	}
	table <- Ball
	<-table
	fmt.Println("球掉了")
}


func player(table chan int,name string){
	for{
		fmt.Printf("%s等待接球\n",name)
		ball := <-table
		fmt.Printf("%s接球%d\n",name,ball)
		ball++
		time.Sleep(1*time.Millisecond)
		fmt.Printf("%s回球%d\n",name,ball)
		table<-ball
	}
}