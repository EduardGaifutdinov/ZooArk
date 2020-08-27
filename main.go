package main

import "https:/src/delivery"

func main() {
	r := delivery.SetupRouter()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
