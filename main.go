package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Product struct {
	Id int
	Name string
	Price float64
	Stock uint
	IsActive bool
}

type API struct {}

var products []Product

func (a *API) GetProducts(empty string, reply *[]Product) error {
	*reply = products

	return nil
}

func (a *API) GetProductById(id int, reply *Product) error {
	var productRes Product

	for _, product := range products {
		if product.Id == id {
			productRes = product
		}
	}

	*reply = productRes

	return nil
}


func (a *API) AddProduct(product Product, reply *Product) error {
	products = append(products, product)
	*reply = product

	return nil
}

func main() {
	products = []Product{
		{
			Id: 0,
			Name: "Phone",
			Price: 150.0,
			Stock: 23,
			IsActive: true,
		},
		{
			Id: 1,
			Name: "Laptop",
			Price: 675.28,
			Stock: 7,
			IsActive: true,
		},
		{
			Id: 2,
			Name: "VGA",
			Price: 1000.13,
			Stock: 10,
			IsActive: false,
		},
	}

	fmt.Println(products)
	
	api := new(API)
	err := rpc.Register(api)
	if (err != nil) {
		log.Fatal("Error register RPC:", err)
	}
	
	rpc.HandleHTTP()
	
	listener, err := net.Listen("tcp", ":4040")
	if (err != nil) {
		log.Fatal("Error listener:", err)
	}
	
	log.Printf("Serving RPC on port %d", 4040)

	err = http.Serve(listener, nil)
	if (err != nil) {
		log.Fatal("Error serving:", err)
	}
}