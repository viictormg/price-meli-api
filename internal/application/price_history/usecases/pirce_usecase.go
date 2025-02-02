package usecases

import (
	"encoding/json"

	"github.com/viictormg/price-meli-api/internal/application/price_history/dto"
	domain "github.com/viictormg/price-meli-api/internal/domain/event"
)

type PriceUsecaseIF interface {
	GetPriceHistory(event domain.Event)
}

type PriceUsecase struct{}

func NewPriceUsecase() PriceUsecaseIF {
	return &PriceUsecase{}
}

func (u *PriceUsecase) GetPriceHistory(event domain.Event) {
	var prices []dto.PriceHistory

	err := json.Unmarshal([]byte(event.Value), &prices)
	if err != nil {
		return
	}

	for _, price := range prices {
		if price.ProductID == "ITEM_ID" {
			continue
		}
	}

}
