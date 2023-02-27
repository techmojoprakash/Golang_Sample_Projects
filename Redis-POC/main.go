package main

import (
	"redispoc/cmd/controller"
)

func main() {
	println("Hello, welcome to Redis-Cache")

	controller.InitializeRedis()
	controller.InitializeRouter()

}
