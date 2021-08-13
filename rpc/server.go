package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	server := rpc.NewServer()
	read := new(XMLReader)
	server.Register(read)
	//rpc.HandleHTTP()

	fmt.Println("rpc server start ...")
	//net.Listen("tcp","")

	http.ListenAndServe("127.0.0.1:3000", server)

}
