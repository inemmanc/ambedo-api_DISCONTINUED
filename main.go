package main

import (
	"ambedo-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	router := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
