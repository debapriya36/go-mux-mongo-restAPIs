package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/debapriya36/mongo-go-mux-crud/routes"
)

func main() {
	fmt.Println("Hey User!🦍")
	r := routes.Router()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
