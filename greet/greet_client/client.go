package main
import (
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
	fmt.Println("Created client: %f", c)
}