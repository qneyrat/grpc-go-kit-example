package client

import (
	"context"
	"log"

	"workshop-go-kit/graphql-gateway/pb"
)

type ProductClient interface {
	GetProduct(id int32) *pb.ProductResponse
	CreateProduct(name string) *pb.ProductResponse
}

type ProductGRPCClient struct {
	Client *GRPCClient
}

func (c *ProductGRPCClient) GetProduct(id int32) *pb.ProductResponse {
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := pb.NewProductClient(conn)
	res, err := client.GetProduct(context.Background(), &pb.GetProductRequest{id})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func (c *ProductGRPCClient) CreateProduct(name string) *pb.ProductResponse {
	conn := c.Client.OpenConn()
	defer conn.Close()

	client := pb.NewProductClient(conn)
	res, err := client.CreateProduct(context.Background(), &pb.CreateProductRequest{name})
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res
}

func NewProductGRPCClient(host string) *ProductGRPCClient {
	return &ProductGRPCClient{NewGRPCClient(host)}
}
