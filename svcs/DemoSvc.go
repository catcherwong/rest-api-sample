package svcs

import (
	"context"
	"log"
)

// DemoSvs is used to implement GreeterServer.
type DemoSvs struct {
	UnimplementedGreeterServer
}

// SayHello implements GreeterServer
func (s *DemoSvs) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}
