package infra

import (
	"github.com/viictormg/price-meli-api/config"
	"github.com/viictormg/price-meli-api/internal/application/price_history/usecases"
	database "github.com/viictormg/price-meli-api/internal/infra/clients"
	"github.com/viictormg/price-meli-api/internal/infra/consumer"
	"github.com/viictormg/price-meli-api/internal/infra/repository"
)

func Run() {
	config := config.NewConfig()
	db := database.NewPostgresConnection()
	priceRepository := repository.NewProductHistoryRepository(db)
	service := usecases.NewPriceUsecase(priceRepository)

	consumer := consumer.NewKafkaConsumer(config.GeKafkaConfg(), service)
	consumer.Consume()
}
