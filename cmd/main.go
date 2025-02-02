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

	
	// // Do something
	// topic := "price"
	// msgCount := 0

	// worker, err := ConnectConsumer([]string{"localhost:9092"})
	// if err != nil {
	// 	panic(err)
	// }

	// consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Connected to Kafka")
	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// doneCh := make(chan struct{})

	// go func() {
	// 	for {
	// 		select {
	// 		case err := <-consumer.Errors():
	// 			fmt.Println(err)
	// 		case msg := <-consumer.Messages():
	// 			msgCount++
	// 			fmt.Println("Recive order cound %d: topic=%s, partition=%d, offset=%d, key=%s, value=%s", msgCount, msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	// 			fmt.Println("Received message", string(msg.Value))
	// 			time.Sleep(200 * time.Millisecond)
	// 		case <-sigchan:
	// 			doneCh <- struct{}{}
	// 		}
	// 	}
	// }()
	// <-doneCh
	// println("Messages processed: ", msgCount)

	// if err := worker.Close(); err != nil {
	// 	panic(err)
	// }

}
