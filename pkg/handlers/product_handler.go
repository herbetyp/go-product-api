package handlers

import (
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func CreateProduct(dto model.ProductDTO) (models.Product, error) {
	prod := models.NewProduct(dto.Name, dto.Price, dto.Code, dto.Qtd, dto.Unity)

	p, err := model.Create(*prod)
	if err != nil {
		return models.Product{}, err
	}
	cacheKeys := []string{utils.PROD_AUTHORIZATION_PREFIX + "all"}
	services.DeleteCache(cacheKeys, false)
	return p, nil
}

func GetProduct(id uint) (models.Product, error) {
	var prod models.Product

	cacheKey := utils.PROD_AUTHORIZATION_PREFIX + utils.UintToString(id)
	if services.GetCache(cacheKey, &prod) == "" {
		p, err := model.Get(id)
		if err != nil {
			return models.Product{}, err
		}
		if p.ID != 0 {
			services.SetCache(cacheKey, &p)
			prod = p
		}

	}
	return prod, nil
}

func GetProducts() ([]models.Product, error) {
	var prods []models.Product

	cacheKey := utils.PROD_AUTHORIZATION_PREFIX + "all"
	if services.GetCache(cacheKey, &prods) == "" {
		ps, err := model.GetAll()
		if err != nil {
			return []models.Product{}, err
		}
		if len(ps) > 0 {
			services.SetCache(cacheKey, &ps)
			prods = ps
		}
	}
	return prods, nil
}

func UpdateProduct(id uint, dto model.ProductDTO) (models.Product, error) {
	prod := models.NewProductWithID(id, dto.Name, dto.Price, dto.Code, dto.Qtd, dto.Unity)

	p, err := model.Update(*prod)
	if err != nil {
		return models.Product{}, err
	}
	cacheKeys := []string{
		utils.PROD_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.PROD_AUTHORIZATION_PREFIX + "all",
	}
	services.DeleteCache(cacheKeys, false)
	return p, nil
}

func DeleteProduct(id uint, hardDelete string) (bool, error) {
	deleted, err := model.Delete(id, hardDelete)

	if err != nil {
		return deleted, err
	}
	cacheKeys := []string{
		utils.PROD_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.PROD_AUTHORIZATION_PREFIX + "all",
	}
	services.DeleteCache(cacheKeys, false)
	return deleted, nil
}

func RecoveryProduct(id uint) (models.Product, error) {
	prod := models.NewProductWithID(id, "", 0, "", 0, "")

	p, err := model.Recovery(*prod)
	if err != nil {
		return models.Product{}, err
	}
	cacheKeys := []string{
		utils.PROD_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.PROD_AUTHORIZATION_PREFIX + "all",
	}
	services.DeleteCache(cacheKeys, false)
	return p, nil
}
