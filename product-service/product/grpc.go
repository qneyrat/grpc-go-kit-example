package product

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"

	"workshop-go-kit/product-service/pb"
)

type GRPCServer struct {
	getProduct    gt.Handler
	createProduct gt.Handler
}

func (s *GRPCServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	//implement
	return nil, nil
}

func (s *GRPCServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	//implement
	return nil, nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.ProductServer {
	return &GRPCServer{
		getProduct: gt.NewServer(
			endpoint.GetProductEndpoint,
			DecodeGetProductRequest,
			EncodeProductResponse,
		),
		createProduct: gt.NewServer(
			endpoint.CreateProductEndpoint,
			DecodeCreateProductRequest,
			EncodeProductResponse,
		),
	}
}
