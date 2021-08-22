package controllers

import (
	"go-web/models"
	"net/http"
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