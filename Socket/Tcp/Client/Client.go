package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:10001")
	checkError(err)
	defer conn.Close()

	conn.Write([]byte("Hello world2!"))
	fmt.Println("send msg")
}

func checkError(err error) {
	if err!=nil{
		fmt.Print("Error:%s",err.Error())
		os.Exit(1)
	}
}
