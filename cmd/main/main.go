package main

import (
	"flexible-payment-go/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Server listening %s\n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
