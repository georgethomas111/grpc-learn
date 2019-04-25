package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/georgethomas111/grpc-learn/learn"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := "0.0.0.0:8080"
	req := &learn.GossipRequest{
		Id: "cli-abc",
	}

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error connecting to grpc server ", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()
	c := learn.NewLearnClient(conn)

	stream, err := c.Gossip(ctx, req)
	if err != nil {
		fmt.Println("Error gossiping ", err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			update, err := stream.Recv()
			if err == io.EOF {
				log.Println("Got eof")
				continue
			}
			if err != nil {
				log.Println("Error receiving msg", err.Error())
			}
			fmt.Println("Got update", update.Info)
		}
	}()

	wg.Wait()
}
