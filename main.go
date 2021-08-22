package main

import (
	"fmt"
	"go-web/routes"
	"net/http"
)

func main() {
	fmt.Println("Running app on http://localhost:8000")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

