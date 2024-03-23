package main

import (
	"ambedo-api/src/config"
	"ambedo-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	api_port := fmt.Sprintf(":%d", config.DefaultApiPort)
	router := router.GenerateRouter()

	fmt.Println("ambedo-API is running on", api_port)
	log.Fatal(http.ListenAndServe(api_port, router))
}
