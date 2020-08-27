package main

import "github.com/ZooArk/src/delivery"

func main() {
	r := delivery.SetupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
