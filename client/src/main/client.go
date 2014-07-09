package main

import (
		"fmt"
		"net"
)

const (
	addr = "127.0.0.1:6666"
)

func main() {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", addr)
	if(err != nil){
		fmt.Println("客户端连接错误：", err.Error())
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	Client(conn)
}

func Client(conn net.Conn) {
	sms := make([]byte, 1024)
	for {
		fmt.Print("请输入要发送的消息:")
		_, err := fmt.Scan(&sms)
		if err != nil {
			fmt.Println("数据输入异常:", err.Error())
		}
		conn.Write(sms)
		buf := make([]byte, 1024)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println(string(buf[0:c]))
		if string(buf[0:c]) == "exit" {
			break
		}
	}

}
