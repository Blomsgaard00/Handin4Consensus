package main

import (
	proto "Handin4Consensus/gRPC"
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	proto.UnimplementedConsensusServer
}

var hasToken bool = true

func (s *Server) HandoverToken(ctx context.Context, token *proto.Token) (*proto.Empty, error) {
	hasToken = true
	return &proto.Empty{}, nil
}

func (s *Server) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5051")
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

	conn, err := grpc.NewClient("localhost:5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client not working")
	}
	//breakpoint wait for terminal
	client := proto.NewConsensusClient(conn)

	client1token := &proto.Token{}
	number := 1
	for true {
		time.Sleep(100 * time.Millisecond)
		if hasToken {
			log.Println("node 1 has token")
			if number%2 == 0 {
				log.Println("node 1 wants to acces cs")
				log.Println("node 1 is inside cs")
				log.Println("node 1 leaves cs")
			}

			client.HandoverToken(context.Background(), client1token)
			hasToken = false
		} else {

			log.Println("node 1 does not have token")
			//waiting for token
		}
		number++
	}
}

func Seed(i int) {
	panic("unimplemented")
}
