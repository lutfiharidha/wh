package handler

import (
	"errors"
	models "ginWeb/app/module/product/model"

	"gorm.io/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Product, error) {
	products := []models.Product{}
	query := db.Preload("Warehouse").Find(&products)
	if err := query.Error; err != nil {
		return products, err
	}

	return products, nil
}

func GetBookByID(id string, db *gorm.DB) (models.Product, bool, error) {
	b := models.Product{}

	query := db.Select("books.*")
	query = query.Group("books.id")
	err := query.Where("books.id = ?", id).First(&b).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteBook(id string, db *gorm.DB) error {
	var b models.Product
	if err := db.Where("id = ? ", id).Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, b *models.Product) error {
	if err := db.Save(&b).Error; err != nil {
		return err
	}
	return nil
}
