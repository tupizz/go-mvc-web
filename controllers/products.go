package controllers

import (
	"go-web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func loadTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func ProductsIndex(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	loadTemplates().ExecuteTemplate(w, "Index", products)
}

func ProductsNew(w http.ResponseWriter, r *http.Request) {
	loadTemplates().ExecuteTemplate(w, "New", nil)
}

func ProductsInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("error on the conversion process")
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("error on the conversion process")
		}

		models.SaveNewProduct(&models.SaveNewProductDTO{
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		})
	}

	http.Redirect(w,r, "/", 301)
}

func ProductsDelete(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query())
	id := r.URL.Query().Get("productId")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}