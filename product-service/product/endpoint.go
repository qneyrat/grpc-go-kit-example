package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type ProductResponse struct {
	ID   int32
	Name string
}

type GetProductRequest struct {
	ID int32
}

type CreateProductRequest struct {
	Name string
}

type Endpoints struct {
	GetProductEndpoint    endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
}

func MakeGetProductEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		gr := req.(GetProductRequest)
		product, _ := svc.GetProduct(gr.ID)
		return ProductResponse{product.ID, product.Name}, nil
	}
}

func MakeCreateProductEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		cr := req.(CreateProductRequest)
		product, _ := svc.CreateProduct(cr.Name)
		return ProductResponse{
			product.ID,
			product.Name,
		}, nil
	}
}
