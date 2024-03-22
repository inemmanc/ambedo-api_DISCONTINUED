package main

import (
	"ambedo-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.GenerateRouter()

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
