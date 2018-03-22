package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vektah/gqlgen/handler"

	"workshop-go-kit/graphql-gateway/client"
	"workshop-go-kit/graphql-gateway/graph"
)

func main() {
	grpcClient := client.NewGRPCClient("product-service:9090")
	productClient := client.ProductGRPCClient{
		grpcClient,
	}

	app := &graph.App{productClient}
	http.Handle("/", handler.Playground("Product", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
