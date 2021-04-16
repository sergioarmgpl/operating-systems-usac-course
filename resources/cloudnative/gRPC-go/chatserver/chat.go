//Package chat definition
package chat

//Imports
import (
	"context"
	"log"
)

//Server struct definition
type Server struct {
}

//SendMessage function implementation
func (s *Server) SendMessage(ctx context.Context, in *MessageRequest) (*MessageReply, error) {
	log.Printf("Received message from: %s", in.GetName())
	return &MessageReply{Body: "Hello " + in.GetName() + " From the Server!"}, nil
}
