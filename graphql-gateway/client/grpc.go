package client

import (
	"log"

	"google.golang.org/grpc"
)

type Client interface {
	OpenConn(host string) *grpc.ClientConn
}

type GRPCClient struct {
	host string
}

func (c *GRPCClient) OpenConn() *grpc.ClientConn {
	cc, err := grpc.Dial(c.host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	return cc
}

func NewGRPCClient(host string) *GRPCClient {
	return &GRPCClient{host}
}
