package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	c := sync.NewCond(&sync.Mutex{})
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		fmt.Println("goroutine1 wait")
		c.Wait()
		fmt.Println("goroutine1")
		c.L.Unlock()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		fmt.Println("goroutine2 wait")
		c.Wait()
		fmt.Println("goroutine2")
		c.L.Unlock()
	}()

	time.Sleep(2*time.Second)
	c.Signal()
	time.Sleep(1*time.Second)
	c.Signal()
	time.Sleep(1*time.Second)
}