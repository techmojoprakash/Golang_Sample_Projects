package main

import (
	"context"
	pc "kafkaapp/internal/kafkapc"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	// fmt.Println("App is invoked")

	// go pc.StartProducer(ctx)
	// fmt.Println("Producer is invoked")
	pc.StartConsumer(ctx)
	// fmt.Println("Consumer is invoked")
	// fmt.Println("Start working with app")
	// time.Sleep(10 * time.Minute)
}
