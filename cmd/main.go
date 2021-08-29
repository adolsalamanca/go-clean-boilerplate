package main

import (
	"fmt"
	"net/http"

	"github.com/adolsalamanca/go-rest-boilerplate/internal"
)

func main() {
	server := internal.NewServer()
	fmt.Printf("Starting server... \n")

	if err := http.ListenAndServe(":8080", server.Routes()); err != nil {
		fmt.Printf("could not initialize server, error :%v \n", err)
	}

}
