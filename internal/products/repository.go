package products

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAll() ([]Product, error) {
	var p []Product

	if err := r.db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (r *Repository) Find(id string) (Product, error) {
	var p Product

	if err := r.db.First(&p, "id = ?", id).Error; err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *Repository) Create(p Product) error {
	return r.db.Create(&p).Error
}

func (r *Repository) Update(p Product) error {
	return r.db.Save(&p).Error
}

func (r *Repository) Delete(p Product) error {
	return r.db.Delete(&p).Error
}
