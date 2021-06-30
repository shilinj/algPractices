package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type (
	Person struct{}

	GreetingRequest struct {
		Name   string
		Gender string
	}

	GreetingResponse struct {
		Greeting string
	}
)

func (p *Person) SayHello(greq *GreetingRequest, gres *GreetingResponse) error {
	hello := fmt.Sprintf("Hello, my name is %v, I'm a %v", greq.Name, greq.Gender)
	gres.Greeting = hello
	return nil
}

func main() {
	rpc.Register(new(Person))
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:3333")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	http.Serve(lis, nil)

}
