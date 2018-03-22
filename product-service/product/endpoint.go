package product

import (
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
	return nil
}

func MakeCreateProductEndpoint(svc Service) endpoint.Endpoint {
	return nil
}
