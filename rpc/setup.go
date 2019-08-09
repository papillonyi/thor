package rpc

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type routeGuideServer struct {
	mu sync.Mutex // protects routeNotes
}

func (s *routeGuideServer) Add(ctx context.Context, in *Point) (*Sum, error) {
	sum := in.One + in.Two + 1
	return &Sum{Sum: sum}, nil
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

func Setup() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 5550))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterRpcServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
