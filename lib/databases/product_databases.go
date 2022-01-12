package databases

import (
	"test_erajaya/config"
	"test_erajaya/models"
)

var res_products []models.ResProducts

func CreateProduct(product *models.Products) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	product_res := models.ResProducts{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	return product_res, nil
}

// function database untuk menampilkan data semua product dari yang terbaru
func GetProductBySortLastCreate() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.created_at DESC").Find(&res_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res_products, nil
}

// function database untuk menampilkan data semua product sesuai A-Z
func GetProductByNameSortAsc() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.name ASC").Find(&res_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res_products, nil
}

// function database untuk menampilkan data semua product sesuai Z-A
func GetProductByNameSortDesc() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.name DESC").Find(&res_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res_products, nil
}

// function database untuk menampilkan data semua product sesuai harga terkecil - terbesar
func GetProductByPriceSortAsc() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.price ASC").Find(&res_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res_products, nil
}

// function database untuk menampilkan data semua product sesuai harga terbesar - terkecil
func GetProductByPriceSortDesc() (interface{}, error) {
	query := config.DB.Table("products").Select("*").Where("products.deleted_at IS NULL").Order("products.price DESC").Find(&res_products)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	return res_products, nil
}
