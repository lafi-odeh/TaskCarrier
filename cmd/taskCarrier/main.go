package main

import (
	"TaskCarrier/api"
	"log"
	"net/http"
)

func main() {
	router := api.SetupRouter()
	log.Println("Starting server on :3000......")
	log.Fatal(http.ListenAndServe(":3000", router))
}
