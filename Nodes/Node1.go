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
	Clients [3]proto.ConsensusClient
	proto.Token
}

func main() {
	connection1, err := grpc.NewClient("localhost:5100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	
}

func (s *Server )HandoverToken(ctx context.Context, token *proto.Token) (*proto.Empty, error){


	return &proto.Empty{}, err
}


func (s *Server)RecieveToken(ctx context.Context, in *Empty) (*Token, error)


