package models

import database "go-web/db"

type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Quantity int
}

func GetAllProducts() []Product {
	db := database.ConnectDb()
	defer db.Close()

	productsSelect, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	var products []Product

	for productsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err!=nil{
			panic(err.Error())
		}

		p := Product{}
		p.Id=id
		p.Name=name
		p.Description=description
		p.Quantity=quantity
		p.Price=price

		products = append(products, p)
	}

	return products
}
