package ports

import "github.com/viictormg/price-meli-api/internal/domain/entity"

type ProductHistoryRepositoryIF interface {
	CreateBulkProductHistory(prices []entity.ProductHistoryEntity) error
}
