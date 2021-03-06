package main
import (
	"context"
	"fmt"
	"log"
	"grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)
func main(){
	fmt.Println("Hello I'm a client")
	conn,err :=grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Println("Created client: %f", c)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient){
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
		FirstName: "One",
		LastName: "Two",
	},
	}
	res,err :=c.Greet(context.Background(), req)
	if err!=nil {
		log.Fatalf("error while calling Greet RPC: %v", err)

	}
    log.Printf("Response from Greet: %v", res.Result)

}

