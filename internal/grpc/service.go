package grpc

import (
	"context"

	"orders/internal/orders"
	pb "orders/proto"
)

type OrderService struct {
	repo orders.Repository
}

func (s *OrderService) ListOrders(
	ctx context.Context,
	req *pb.ListOrdersRequest,
) (*pb.ListOrdersResponse, error) {
	orders, err := s.repo.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	var response pb.ListOrdersResponse
	for _, order := range orders {
		response.Orders = append(response.Orders, &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			OrderDate:    order.OrderDate,
			TotalAmount:  order.TotalAmount,
		})
	}

	return &response, nil
}
