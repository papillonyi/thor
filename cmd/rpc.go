package main

import (
	"github.com/papillonyi/thor/pkg/setting"
	"github.com/papillonyi/thor/rpc"
)

func init() {
	setting.Setup()
}

func main() {
	rpc.Setup()
}
