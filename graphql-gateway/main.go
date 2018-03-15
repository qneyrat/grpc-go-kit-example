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
	clt := client.NewProductGRPCClient("product-service:9090")
	app := &graph.App{clt}
	http.Handle("/", handler.Playground("Product", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
