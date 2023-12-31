package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"cartify/order/infrastructure/repository"
	"cartify/order/interfaces"

	amqp "github.com/rabbitmq/amqp091-go"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()

	uri := os.Getenv("DATABASE")
	if uri == "" {
		uri = "mongodb://localhost:27017/orders"
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	rabbitmqUri := os.Getenv("RABBITMQ")
	if rabbitmqUri == "" {
		rabbitmqUri = "amqp://guest:guest@localhost:5672/"
	}
	conn, err := amqp.Dial(rabbitmqUri)
	if err != nil {
		log.Print("Failed to connect to RabbitMQ")
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Print("Failed to open a channel")
		panic(err)
	}
	defer ch.Close()

	orderRepository := repository.NewOrderRepository(client)
	eventRepository := repository.NewEventRepository(ch)

	server := interfaces.NewServer(interfaces.ServerParams{
		OrderRepository: orderRepository,
		EventRepository: eventRepository,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("server listening at", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
