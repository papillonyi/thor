package main

import (
	"github.com/papillonyi/thor/api"
	"github.com/papillonyi/thor/model"
)

func main() {
	model.Migrate()
	api.Load()
}
