package main

import (
	"github.com/IBM/sarama"
	"github.com/viictormg/price-meli-api/internal/infra"
)

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)

}
func main() {
	infra.Run()
}
