package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/georgethomas111/grpc-learn/learn"
	"google.golang.org/grpc"
)

type StreamingServer struct {
	streams map[string]learn.Learn_GossipServer
}

func (s *StreamingServer) Gossip(req *learn.GossipRequest, stream learn.Learn_GossipServer) error {
	if req == nil {
		return errors.New("Request cannot be nil")
	}
	s.streams[req.Id] = stream
	return nil
}

// Update time updates the current time to all the clients.
func (s *StreamingServer) UpdateTime() {
	msg := &learn.Messages{
		Info: strconv.Itoa(int(time.Now().Unix())),
	}

	fmt.Println("Time to update time to clients ", len(s.streams))
	for id, stream := range s.streams {
		err := stream.Send(msg)
		if err != nil {
			fmt.Println("Error streaming message from server to client with id ", id)
		}
	}
}

func (s *StreamingServer) Start(grpcAddr string) error {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	g := grpc.NewServer()
	learn.RegisterLearnServer(g, s)

	fmt.Println("Listening for grpc requests on ", grpcAddr)
	return g.Serve(lis)
}

func New() *StreamingServer {
	return &StreamingServer{
		streams: make(map[string]learn.Learn_GossipServer),
	}
}

func main() {
	grpcAddr := "0.0.0.0:8080"
	s := New()
	go func() {
		for {
			<-time.After(time.Second)
			s.UpdateTime()
		}
	}()
	err := s.Start(grpcAddr)
	if err != nil {
		fmt.Println("Error starting server ", err.Error())
	}
}
