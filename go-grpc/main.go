package main

import (
	"fmt"
	"grpc_tutorial/grpc"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := grpc.Person{
		Id:    1234,
		Name:  "John Wick",
		Email: "jwick@continental.com",
		Phones: []*grpc.Person_PhoneNumber{
			{Number: "777-777-7777", Type: grpc.Person_HOME},
		},
	}

	out, err := proto.Marshal(&p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(len(out))
}
