package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vektah/gqlgen/handler"

	"grpc-go-kit-example/graphql-gateway/client"
	"grpc-go-kit-example/graphql-gateway/graph"
)

func main() {
	clt := client.NewProductGRPCClient("product-service:9090")
	app := &graph.App{clt}
	http.Handle("/", handler.Playground("Product", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("HTTP listen on 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
