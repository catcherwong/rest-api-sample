package svcs

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"net"
)

func InitGRPC() {

	lis, err := net.Listen("tcp", ":19999")
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()

	RegisterGreeterServer(server, &DemoSvs{})

	log.Println("start grpc service in 19999")

	server.Serve(lis)

}
