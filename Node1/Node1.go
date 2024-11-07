package main

import (
	proto "Handin4Consensus/gRPC"
	"sync"
	"context"
	"fmt"
	"net"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	proto.UnimplementedConsensusServer
}

func (s *Server )HandoverToken(ctx context.Context, token *proto.Token) (*proto.Empty, error){
	

	return &proto.Empty{}, nil
}


func (s *Server) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
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
	server.start_server()
	
	conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client not working")
	}
	//breakpoint wait for terminal 
	client := proto.NewConsensusClient(conn)

	client1token := &proto.Token{
		Token: true,
	}
	
	for true{
		if(client1token.Token){
			client.HandoverToken(context.Background(), client1token)


			client1token.Token = false 
		}else 
		{
			//waiting for token
		}
		
		//recieving token
		//entering critical section
		//leaving critcial section
	
	}
}

