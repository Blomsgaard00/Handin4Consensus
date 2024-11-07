package main

import (
	proto "Handin4Consensus/gRPC"
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Server struct {
	proto.UnimplementedConsensusServer
}

var hasToken bool = false

func (s *Server) HandoverToken(ctx context.Context, token *proto.Token) (*proto.Empty, error) {
	hasToken = true
	return &proto.Empty{}, nil
}

func (s *Server) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5053")
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterConsensusServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Did not work")
	}
}

func main() {
	server := Server{}
	go server.start_server()

	conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client not working")
	}
	//breakpoint wait for terminal
	client := proto.NewConsensusClient(conn)

	client1token := &proto.Token{
		Token: false,
	}

	for true {
		time.Sleep(1000)
		if hasToken {
			fmt.Println("node 3 has token")
			client.HandoverToken(context.Background(), client1token)
			hasToken = false
		} else {
			//fmt.Println("node 3 does not have token")
			//waiting for token
		}

		//recieving token
		//entering critical section
		//leaving critcial section

	}
}
