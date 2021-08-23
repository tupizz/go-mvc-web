package main

import (
	"go-web/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Running app on http://localhost:8000")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

