package order

import (
	"cartify/order/domain/model"
	"cartify/order/domain/repository"
	"context"
)

type ListOrdersParams struct {
	CustomerId string
}

type ListOrders func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error)

func NewListOrders(orderRepository repository.OrderRepository) ListOrders {
	return func(ctx context.Context, params ListOrdersParams) ([]*model.Order, error) {
		return orderRepository.ListOrders(ctx, params.CustomerId)
	}
}
