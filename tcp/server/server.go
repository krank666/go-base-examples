package main

import (
	"bufio"
	"fmt"
	"net"
)

func main()  {
	tcpAddr,_ := net.ResolveTCPAddr("tcp",":9999")
	tcpListen,_ := net.ListenTCP("tcp",tcpAddr)
	for{
		conn,err :=	tcpListen.AcceptTCP()
		if err!=nil{
			fmt.Println(err.Error())
			return
		}
		go handleTcp(conn)
	}
}


func handleTcp(conn *net.TCPConn){
	reader := bufio.NewReader(conn)
	for{
		bufBy := make([]byte,1024)
		n,_ := reader.Read(bufBy)
		fmt.Println("收到了"+conn.RemoteAddr().String()+"丢过来的"+string(bufBy[0:n]))
		conn.Write(bufBy[0:n])
		fmt.Println("然后给他丢回去了")
	}
	defer conn.Close()
}