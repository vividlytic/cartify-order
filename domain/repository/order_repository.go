package repository

import (
	"context"

	"cartify/order/domain/model"
)

type OrderRepository interface {
	ListOrders(ctx context.Context, customerId string) ([]*model.Order, error)
	GetOrder(ctx context.Context, id string) (*model.Order, error)
	CreateOrder(ctx context.Context, customerId string, customerName string, orderItem []model.OrderItem) (string, error)
}
