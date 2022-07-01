package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"github.com/wonpanu/learn-golang/amqputil"
)

func main() {
	fmt.Println("Bulk worker started!")
	godotenv.Load(".env")

	amqpHost := os.Getenv("AMQP_HOST")
	amqpPortRaw := os.Getenv("AMQP_PORT")
	amqpUsername := os.Getenv("AMQP_USERNAME")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	blogQueueName := os.Getenv("LOG_API_QUEUE_NAME")

	amqpPort, err := strconv.Atoi(amqpPortRaw)
	if err != nil {
		log.Fatal(err)
	}

	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", amqpUsername, amqpPassword, amqpHost, amqpPort)
	fmt.Println("rabbitMQ URI:", amqpURI)

	amqpConn, amqpCh, _, msgs := amqputil.CreateConsumerConnection(amqpURI, blogQueueName)
	amqpCloseNotify := amqpConn.NotifyClose(make(chan *amqp.Error))
	defer func() {
		amqpCh.Close()
		amqpConn.Close()
	}()

	go func() {
		for err := range amqpCloseNotify {
			log.Println("Rabbit MQ connection lost", err)
			os.Exit(1)
		}
	}()

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
