package rpc

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

func _add(client RpcClient, one int32, two int32) int32 {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sum, err := client.Add(ctx, &Point{One: one, Two: two})
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return sum.Sum
}

func CallAdd(one int32, two int32) int32 {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("10.4.145.242:5550", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewRpcClient(conn)
	return _add(client, one, two)

}
