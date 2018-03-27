package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"grpc-go-kit-example/product-service/database"
	"grpc-go-kit-example/product-service/pb"
	"grpc-go-kit-example/product-service/product"
)

func main() {
	ctx := context.Background()
	productService := &product.ProductService{database.NewDBProductRepository()}
	errors := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := grpc.NewServer()
		pb.RegisterProductServer(gRPCServer, product.NewGRPCServer(ctx, product.Endpoints{
			GetProductEndpoint:    product.MakeGetProductEndpoint(productService),
			CreateProductEndpoint: product.MakeCreateProductEndpoint(productService),
		}))

		fmt.Println("gRPC listen on 9090")
		errors <- gRPCServer.Serve(listener)
	}()

	fmt.Println(<-errors)
}
