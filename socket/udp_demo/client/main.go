package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp  client

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10808,
	})
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err := c.Write([]byte(s))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			return
		}
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
		}
		fmt.Printf("read from %v,mag:%v\n", addr, string(buf[:n]))
	}
}
