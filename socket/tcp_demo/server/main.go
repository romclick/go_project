package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server
func process(conn net.Conn) {
	defer conn.Close() //处理完要关闭连接
	//针对当前链接做数据的发送和接收
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte

		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("recv:", recv)
		conn.Write([]byte("okk")) //把收到的数据返回
	}

}

func main() {
	//1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:10808")
	if err != nil {
		fmt.Printf("listen failed ,err:%v\n", err)
	}
	for {
		//2.等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed ,err:%v\n", err)
		}
		//3.启动一个单独的goroutine来处理
		go process(conn)
	}
}
