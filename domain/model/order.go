package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID           primitive.ObjectID `bson:"_id"`
	CustomerId   string             `bson:"customerId"`
	CustomerName string             `bson:"customerName"`
	OrderItem    []OrderItem        `bson:"orderItem"`
}

type OrderItem struct {
	ItemId    int `bson:"itemId"`
	Quantity  int `bson:"quantity"`
	UnitPrice int `bson:"unitPrice"`
}

type OrderEvent struct {
	ID           string
	CustomerId   string
	CustomerName string
	OrderItem    []OrderItem
}
