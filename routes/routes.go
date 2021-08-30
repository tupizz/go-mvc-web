package routes

import (
	"go-web/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.ProductsIndex)
	http.HandleFunc("/new", controllers.ProductsNew)
	http.HandleFunc("/insert", controllers.ProductsInsert)
	http.HandleFunc("/delete", controllers.ProductsDelete)
}