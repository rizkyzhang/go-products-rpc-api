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
			break
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

func (a *API) UpdateProductById(updatedProduct Product, reply *Product) error {
	var newProducts []Product

	for _, product := range products {
		if product.Id == updatedProduct.Id {
			newProducts = append(newProducts, updatedProduct) 
		} else {
			newProducts = append(newProducts, product)
		}
	}

	products = newProducts
	*reply = updatedProduct

	return nil
}

func (a *API) DeleteProductById(id int, reply *Product) (error) {
	var deletedProduct Product

	for _, product := range products {
		if product.Id == id {
			products = append(products[:id], products[id + 1:]...)
			deletedProduct = product
			break
		}
	}

	*reply = deletedProduct

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