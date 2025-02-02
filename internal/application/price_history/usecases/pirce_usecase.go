package usecases

import (
	"encoding/json"
	"time"

	"github.com/viictormg/price-meli-api/internal/application/price_history/dto"
	"github.com/viictormg/price-meli-api/internal/application/price_history/ports"
	"github.com/viictormg/price-meli-api/internal/domain/entity"
	domain "github.com/viictormg/price-meli-api/internal/domain/event"
)

type PriceUsecaseIF interface {
	GetPriceHistory(event domain.Event)
}

type PriceUsecase struct {
	repository ports.ProductHistoryRepositoryIF
}

func NewPriceUsecase(repository ports.ProductHistoryRepositoryIF) PriceUsecaseIF {
	return &PriceUsecase{
		repository: repository,
	}
}

func (u *PriceUsecase) GetPriceHistory(event domain.Event) {
	var prices []dto.PriceHistory
	var pricesToInsert []entity.ProductHistoryEntity

	err := json.Unmarshal([]byte(event.Value), &prices)
	if err != nil {
		return
	}

	for _, price := range prices {
		if price.ProductID == "ITEM_ID" {
			continue
		}

		newPrice := entity.NewProductHistoryEntity(price.ProductID, price.Price, price.OrderDate)
		pricesToInsert = append(pricesToInsert, newPrice)
	}

	u.repository.CreateBulkProductHistory(pricesToInsert)
	time.Sleep(500 * time.Millisecond)
}
