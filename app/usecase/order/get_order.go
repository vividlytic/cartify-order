package order

import (
	"cartify/order/domain/model"
	"cartify/order/domain/repository"
	"context"
)

type GetOrdersParams struct {
	ID string
}

type GetOrder func(ctx context.Context, params GetOrdersParams) (*model.Order, error)

func NewGetOrder(orderRepository repository.OrderRepository) GetOrder {
	return func(ctx context.Context, params GetOrdersParams) (*model.Order, error) {
		return orderRepository.GetOrder(ctx, params.ID)
	}
}
