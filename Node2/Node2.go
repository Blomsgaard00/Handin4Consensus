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
	Clients [2]proto.ConsensusClient
	proto.Token
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
	server.start_server()
	
}



func (s *Server )HandoverToken(ctx context.Context, token *proto.Token) (*proto.Empty, error){


	return &proto.Empty{}, err
}


func (s *Server)RecieveToken(ctx context.Context, in *Empty) (*Token, error)

