package main

import (
	"github.com/papillonyi/thor/mq"
	"github.com/papillonyi/thor/pkg/setting"
)

func main() {
	setting.Setup()
	mq.SetupWork()
}
