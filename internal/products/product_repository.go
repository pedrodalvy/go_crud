package products

import (
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	var p []Product

	if err := r.db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (r *ProductRepository) Find(id string) (Product, error) {
	var p Product

	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *ProductRepository) Create(p Product) error {
	return r.db.Create(&p).Error
}

func (r *ProductRepository) Update(p Product) error {
	return r.db.Save(&p).Error
}

func (r *ProductRepository) Delete(id string) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}
