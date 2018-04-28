package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	udp_addr,err:=net.ResolveUDPAddr("udp",":11110")
	checkError(err)
	conn,err:=net.ListenUDP("udp",udp_addr)
	defer conn.Close()
	checkError(err)
	recvUDPMsg(conn)
}

func recvUDPMsg(conn *net.UDPConn) {
	var buf [20]byte

	for {
		n, raddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("msg is:", string(buf[0:n]))
		_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
		checkError(err)
	}
}

func checkError(err error) {
	if err!=nil{
		fmt.Println("Error:%s",err.Error())
		os.Exit(1)
	}
}
