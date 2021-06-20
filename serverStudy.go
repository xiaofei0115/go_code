package main

import (
	"fmt"
	"net"
)

func main() {
	//打开连接
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listener err=", err)
	}
	defer listener.Close()
	//阻塞连接 等待输入
	conn, err := listener.Accept()
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()
	//读取数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("#%v#", string(buf[:n]))
}
