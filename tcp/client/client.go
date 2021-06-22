package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	tcpAddr,_ := net.ResolveTCPAddr("tcp",":9999")
	dialTCP,_ := net.DialTCP("tcp",nil,tcpAddr)

	reader:= bufio.NewReader(os.Stdin)
	for{
		bufBy:=make([]byte,1024)
		str,_:= reader.ReadString('\n')
		newStr := strings.Replace(str, "\n", "", -1)
		dialTCP.Write([]byte(newStr))
		n,_:= dialTCP.Read(bufBy)
		fmt.Println("收到了丢回来的"+string(bufBy[0:n]))
	}
}
