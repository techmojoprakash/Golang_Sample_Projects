package main

import (
	"compavg/CompApp/comppb"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8082"
)

func doClientStream(c comppb.CompServiceClient) {
	fmt.Println("doClientStream Func init...!")
	requests := []*comppb.CompRequest{
		{
			Num: 10,
		},
		{
			Num: 11,
		}, {
			Num: 12,
		}, {
			Num: 13,
		}, {
			Num: 14,
		}, {
			Num: 15,
		}, {
			Num: 16,
		}, {
			Num: 17,
		}, {
			Num: 18,
		},
	}

	stream, err := c.CompAvg(context.Background())
	if err != nil {
		log.Fatal("Error while calling doClientStream ", err)
	}

	// loop over individual items in requeests slice and send to server
	for _, req := range requests {
		fmt.Println("Sending req", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receving response from CompAvg ", err)
	}
	fmt.Println("CompAvg Response is ", resp)
}

func main() {
	fmt.Println("Client init......")
	// cc = Client Connection
	cc, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client Could not connect to gRPC Server : %v \n", err)
	}
	defer cc.Close()
	conn := comppb.NewCompServiceClient(cc)
	fmt.Printf("Client created successfully and connected port is, %v \n", PORT)

	// Call Client Stream func
	doClientStream(conn)
}
