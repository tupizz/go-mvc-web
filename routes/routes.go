package routes

import (
	"go-web/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.ProductsIndex)
	http.HandleFunc("/new", controllers.ProductsNew)
}