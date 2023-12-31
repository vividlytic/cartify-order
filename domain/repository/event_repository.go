package repository

import (
	"context"

	"cartify/order/domain/model"
)

type EventRepository interface {
	PostOrderEvent(ctx context.Context, orderEvent model.OrderEvent) error
}
