package main

import (
	"log"
	"net/http"
	"backend/api/routes"
)

func main() {
	routes.InitializeRoutes()

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
