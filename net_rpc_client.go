package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type GreetingRequest struct {
	Name   string
	Gender string
}

type GreetingResponse struct {
	Greeting string
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:3333")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := GreetingRequest{"LiLei", "boy"}
	var res GreetingResponse

	err = conn.Call("Person.SayHello", &req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Println(res.Greeting)
}
