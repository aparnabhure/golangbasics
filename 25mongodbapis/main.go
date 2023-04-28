package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aparnabhure/golangmongodbapis/router"
)

// Mongo GO lib https://github.com/mongodb/mongo-go-driver
func main() {
	fmt.Println("Welcom to Mongo DB APIs")

	router := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening on port 4000...")
}
