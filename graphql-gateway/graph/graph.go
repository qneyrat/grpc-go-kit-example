package graph

import (
	"context"

	"grpc-go-kit-example/graphql-gateway/client"
)

type App struct {
	ProductClient *client.ProductGRPCClient
}

func (a *App) Query_product(ctx context.Context, id int) (Product, error) {
	res := a.ProductClient.GetProduct(int32(id))
	return Product{
		int(res.Id),
		res.Name,
	}, nil
}

func (a *App) Mutation_createProduct(ctx context.Context, name string) (Product, error) {
	res := a.ProductClient.CreateProduct(name)
	return Product{
		int(res.Id),
		res.Name,
	}, nil
}
