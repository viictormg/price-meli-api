package consumer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/viictormg/price-meli-api/config"
	"github.com/viictormg/price-meli-api/internal/application/price_history/usecases"
	domain "github.com/viictormg/price-meli-api/internal/domain/event"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
	topic    string
	service  usecases.PriceUsecaseIF
}

func NewKafkaConsumer(
	config *config.KafkaConfig,
	service usecases.PriceUsecaseIF,
) *KafkaConsumer {

	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Errors = true

	consumer, err := sarama.NewConsumer(config.Brokers, saramaConfig)

	if err != nil {
		panic(err)
	}

	return &KafkaConsumer{
		consumer: consumer,
		topic:    strings.TrimSpace(config.Topic),
		service:  service,
	}
}

func (k *KafkaConsumer) Consume() {
	partitions, err := k.consumer.Partitions(k.topic)

	if err != nil {
		fmt.Println("Error getting partitions: ", err)
	}

	fmt.Println("Connected to Kafka")
	// Manejo de señales para salida segura
	ctx, cancel := context.WithCancel(context.Background())
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		fmt.Println("\nInterrupt signal received. Closing consumers...")
		cancel()
	}()

	for _, partition := range partitions {
		pc, err := k.consumer.ConsumePartition(k.topic, partition, sarama.OffsetNewest)

		if err != nil {
			fmt.Println("Error consuming partition: ", err)
		}

		go func(pc sarama.PartitionConsumer) {
			for {

				select {
				case err := <-pc.Errors():
					fmt.Println("Error consuming message: ", err)
				case msg := <-pc.Messages():
					event := domain.Event{
						Key:   string(msg.Key),
						Value: string(msg.Value),
					}

					k.service.GetPriceHistory(event)
				}
			}
		}(pc)
	}

	<-ctx.Done()
	println("Messages processed: ")

	k.Close()
}

func (k *KafkaConsumer) Close() {
	k.consumer.Close()
}
