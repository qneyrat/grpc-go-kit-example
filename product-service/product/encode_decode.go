package product

import (
	"context"

	"workshop-go-kit/product-service/pb"
)

func EncodeProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ProductResponse)
	return &pb.ProductResponse{
		res.ID,
		res.Name,
	}, nil
}

func DecodeGetProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetProductRequest)
	return GetProductRequest{req.Id}, nil
}

func DecodeCreateProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateProductRequest)
	return CreateProductRequest{req.Name}, nil
}
