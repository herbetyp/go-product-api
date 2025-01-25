package handlers

import (
	"log"

	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
)

func CreateProduct(dto model.ProductDTO) (models.Product, error) {
	product := models.NewProduct(dto.Name, dto.Price, dto.Code, dto.Qtd, dto.Unity)

	p, err := model.Create(*product)
	if err != nil {
		log.Printf("cannot create product: %v", err)
		return models.Product{}, err
	}
	return p, nil
}

func GetProduct(id uint) (models.Product, error) {
	p, err := model.Get(id)

	if err != nil {
		log.Printf("cannot find product: %v", err)
		return models.Product{}, err
	}
	return p, nil
}

func GetProducts() ([]models.Product, error) {
	ps, err := model.GetAll()

	if err != nil {
		log.Printf("error on get products: %v", err)
		return []models.Product{}, err
	}
	return ps, nil
}
