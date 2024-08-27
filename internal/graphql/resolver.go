package graphql

import (
	"context"

	"orders/internal/orders"
)

type Resolver struct {
	repo orders.Repository
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*Order, error) {
	return r.repo.ListOrders(ctx)
}
