package databases

import (
	"test_erajaya/config"
	"test_erajaya/models"
)

func CreateProduct(product *models.Products) (interface{}, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
