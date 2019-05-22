package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-go/greet/greetpb"




)

type server struct{

}

func (*server) Greet (ctx context.Context, req *greetpb.GreetRequest )(*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v",req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res,nil

}





func main(){
	fmt.Println("Hello World")

    //lis,err := net.Listen(network string, dddress string)
	lis,err := net.Listen("tcp", "0.0.0.0:5051")
	if err!=nil {
		log.Fatalf("Failed to listen: %v",err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis) ; err!=nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}