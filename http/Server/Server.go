package main

import (
	"net/http"
	"log"
	"io"
)

func main(){
	http.HandleFunc("/add",AddHandler)
	http.HandleFunc("/",IndexHandler)
	err:=http.ListenAndServe(":9091",nil)
	if err!=nil{
		log.Fatal("AddHandler",err)
	}
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"index")
}
func AddHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"add\n")
}
