//Package main definition
package main

//Imports
import (
	"fmt"
	"log"
	"net"

	chat "grpcTutorial/chatserver"

	"google.golang.org/grpc"
)

//Main definition
func main() {
	println("gRPC Server")

	//Listening port definition, in this case port 9000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	println("gRPC Server now listening on port 9000")

	//New instance or the gRPC server
	grpcServer := grpc.NewServer()

	//Registrtion of the service with the gRPC server
	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	//We call Serve() to start the service
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
