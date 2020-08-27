package main

import (
	"github.com/ZooArk/src/delivery"
	_ "github.com/ZooArk/docs"
)

// @title ZooArk
// @version 1.0.0
func main() {
	r := delivery.SetupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
