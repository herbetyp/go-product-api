package handlers

import (
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
)

func CreateProduct(dto model.ProductDTO) (models.Product, error) {
	prod := models.NewProduct(dto.Name, dto.Price, dto.Code, dto.Qtd, dto.Unity)

	p, err := model.Create(*prod)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func GetProduct(id uint) (models.Product, error) {
	p, err := model.Get(id)

	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func GetProducts() ([]models.Product, error) {
	ps, err := model.GetAll()

	if err != nil {
		return []models.Product{}, err
	}
	return ps, nil
}

func UpdateProduct(id uint, dto model.ProductDTO) (models.Product, error) {
	prod := models.NewProductWithID(id, dto.Name, dto.Price, dto.Code, dto.Qtd, dto.Unity)

	p, err := model.Update(*prod)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func DeleteProduct(id uint, hardDelete string) (bool, error) {
	deleted, err := model.Delete(id, hardDelete)

	if err != nil {
		return deleted, err
	}
	return deleted, nil
}

func RecoveryProduct(id uint) (models.Product, error) {
	prod := models.NewProductWithID(id, "", 0, "", 0, "")

	p, err := model.Recovery(*prod)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}
