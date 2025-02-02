package infra

import (
	"github.com/viictormg/price-meli-api/config"
	"github.com/viictormg/price-meli-api/internal/application/price_history/usecases"
	"github.com/viictormg/price-meli-api/internal/infra/consumer"
)

func Run() {
	config := config.NewConfig()
	service := usecases.NewPriceUsecase()

	consumer := consumer.NewKafkaConsumer(config.GeKafkaConfg(), service)
	consumer.Consume()
}
