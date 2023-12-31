package order

import (
	"cartify/order/domain/model"
	"cartify/order/domain/repository"
	"context"
)

type CreateOrderParams struct {
	CustomerId   string
	CustomerName string
	OrderItem    []model.OrderItem
}

type CreateOrder func(ctx context.Context, params CreateOrderParams) (string, error)

func NewCreateOrder(orderRepository repository.OrderRepository, eventRepository repository.EventRepository) CreateOrder {
	return func(ctx context.Context, params CreateOrderParams) (string, error) {
		orderId, err := orderRepository.CreateOrder(ctx, params.CustomerId, params.CustomerName, params.OrderItem)
		if err != nil {
			return "", err
		}

		event := model.OrderEvent{
			ID:           orderId,
			CustomerId:   params.CustomerId,
			CustomerName: params.CustomerName,
			OrderItem:    params.OrderItem,
		}

		err = eventRepository.PostOrderEvent(ctx, event)
		return orderId, err
	}
}
