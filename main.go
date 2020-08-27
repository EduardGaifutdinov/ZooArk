package main

import (
	_ "github.com/ZooArk/docs"
	"github.com/ZooArk/src/delivery"
)

// @title ZooArk
// @version 1.0.0
func main() {
	r := delivery.SetupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
