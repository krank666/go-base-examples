package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f,err:= os.OpenFile("./test.txt",os.O_CREATE|os.O_RDWR ,0666)
	f2,err2:= os.OpenFile("./testCopy.txt",os.O_CREATE|os.O_RDWR ,0666)
	fmt.Println(err,err2)
	io.Copy(f2,f)
}

func read(s int){
	fmt.Println(s)
}
