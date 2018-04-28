package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	listen_socket,err:=net.Listen("tcp","localhost:10001")
	checkError(err)
	fmt.Println("liesten at :10001")
	defer listen_socket.Close()

	for{
		new_conn,err:=listen_socket.Accept()
		if err !=nil{
			continue
		}
		go recvConnMsg(new_conn)
	}
}
func recvConnMsg(conn net.Conn) {
	buf:=make([]byte,50)
	defer conn.Close()
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("conn closed")
			return
		}
		fmt.Println("recv smg:",string(buf[0:n]))
	}
}
func checkError(err error) {
	if err!=nil{
		fmt.Println("Error:%s",err.Error())
		os.Exit(1)
	}
}
