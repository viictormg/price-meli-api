package repository

import (
	"github.com/viictormg/price-meli-api/internal/application/price_history/ports"
	"github.com/viictormg/price-meli-api/internal/domain/entity"
	"gorm.io/gorm"
)

type ProductHistoryRepository struct {
	db *gorm.DB
}

func NewProductHistoryRepository(db *gorm.DB) ports.ProductHistoryRepositoryIF {
	return &ProductHistoryRepository{db}
}

func (r *ProductHistoryRepository) CreateBulkProductHistory(prices []entity.ProductHistoryEntity) error {
	return r.db.Create(&prices).Error
}
