package main

import (
	"fmt"
	"go_ecommerce-app/server"
)

func main() {

	fmt.Println("Starting the server...")
	fmt.Println("Server is running on port 3000")
	server.StartServer()

}
