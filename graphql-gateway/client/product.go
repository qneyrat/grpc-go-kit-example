package client

import (
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
	//implement
	return &pb.ProductResponse{0, "not implemented"}
}

func (c *ProductGRPCClient) CreateProduct(name string) *pb.ProductResponse {
	//implement
	return &pb.ProductResponse{0, name}
}

func NewProductGRPCClient(host string) *ProductGRPCClient {
	return &ProductGRPCClient{NewGRPCClient(host)}
}
