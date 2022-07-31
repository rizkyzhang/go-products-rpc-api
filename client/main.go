package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Product struct {
	Name string
	Id int
	Price float64
	Stock uint
	IsActive bool
}

func main() {
	var products []Product
	var product Product

	client, err := rpc.DialHTTP("tcp", ":4040")
	if (err != nil) {
		log.Fatal("RPC Dial error:", err)
	}

	client.Call("API.GetProductById", 1, &product)
	fmt.Println("Get product by id:", product)
	fmt.Println("")

	client.Call("API.AddProduct", Product{
			Id: 3,
			Name: "Drone",
			Price: 999.0,
			Stock: 29,
			IsActive: false,
		}, &product)
	fmt.Println("Add product:", product)
	fmt.Println("")

	client.Call("API.UpdateProductById", Product{
			Id: 3,
			Name: "Robot",
			Price: 235.0,
			Stock: 9,
			IsActive: true,
		}, &product)
	fmt.Println("Update product by id:", product)
	fmt.Println("")

	client.Call("API.DeleteProductById", 3, &product)
	fmt.Println("Delete product by id:", product)
	fmt.Println("")

	client.Call("API.GetProducts", "", &products)
	fmt.Println("Get products:", products)
	fmt.Println("")
}
