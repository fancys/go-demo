package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	http, err := rpc.DialHTTP("tcp", "127.0.0.1:3000")
	fmt.Println("remote calling ..")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	input := Input{"hello.txt"}
	var output OutPut
	//http.Go()
	err = http.Call("XMLReader.Read", input, &output)
	if err != nil {
		fmt.Println("remote call error:", err)
		panic(err)
	}
	fmt.Println(output.FileName)
}
