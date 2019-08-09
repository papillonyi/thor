package main

import (
	"github.com/papillonyi/thor/pkg/setting"
	"github.com/papillonyi/thor/rpc"
	"log"
)

func init() {
	setting.Setup()
}

func main() {
	log.Printf("get sum: %d", 1)
	for i := 1; i <= 10000; i++ {
		rpc.CallAdd(int32(i), 2)
	}
	log.Printf("get sum: %d", 2)
}
