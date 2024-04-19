package healthcheck

import (
	"fmt"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedHealthCheckServer
}

func (s *Server) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	fmt.Println("Received ping request")
	return &PingResponse{Status: "200", Message: "Ok"}, nil
}
