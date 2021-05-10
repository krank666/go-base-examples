package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	file,err := os.Open("README.MD")
	if(err!=nil){
		fmt.Println(err)
	}else{
		inputReader := bufio.NewReader(file)
		var Count int = 0
		for {
		     Count += 1
			    inputString, readerError := inputReader.ReadString('\n')  //我们将inputReader里面的字符串按行进行读取。
		    if readerError == io.EOF {
			            return  //如果遇到错误就终止循环。
			       }
			        fmt.Printf("The %d line is: %s",Count, inputString)    //将文件的内容逐行（行结束符'\n'）读取出来。
		    }
	}
	defer file.Close()
}
