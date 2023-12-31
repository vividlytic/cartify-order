package service

import (
	usecase "cartify/order/app/usecase/order"
	"context"

	"cartify/order/domain/model"
	pb "cartify/order/proto/order"
)

type OrderServiceServer struct {
	listOrders  usecase.ListOrders
	getOrder    usecase.GetOrder
	createOrder usecase.CreateOrder
	pb.UnimplementedOrderServiceServer
}

func NewOrderServer(
	listOrders usecase.ListOrders,
	getOrder usecase.GetOrder,
	createOrder usecase.CreateOrder,
) pb.OrderServiceServer {
	return &OrderServiceServer{
		listOrders:  listOrders,
		getOrder:    getOrder,
		createOrder: createOrder,
	}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	// if md, ok := metadata.FromIncomingContext(ctx); ok {
	// 	fmt.Println(md.Get("authorization"))
	// }
	params := usecase.ListOrdersParams{CustomerId: request.GetCustomerId()}
	orders, err := s.listOrders(ctx, params)
	if err != nil {
		return nil, err
	}

	protoOrders := make([]*pb.Order, 0)

	for _, b := range orders {
		protoOrders = append(protoOrders, OrderToProto(b))
	}

	response := &pb.ListOrdersResponse{Orders: protoOrders}
	return response, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	params := usecase.GetOrdersParams{ID: request.OrderId}

	order, err := s.getOrder(ctx, params)
	if err != nil {
		return nil, err
	}

	protoOrder := OrderToProto(order)

	return &pb.GetOrderResponse{Order: protoOrder}, nil
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	protoOrderItem := request.GetOrderItem()
	orderItem := make([]model.OrderItem, len(protoOrderItem))

	for i, v := range protoOrderItem {
		orderItem[i].ItemId = int(v.ItemId)
		orderItem[i].Quantity = int(v.Quantity)
		orderItem[i].UnitPrice = int(v.UnitPrice)
	}
	params := usecase.CreateOrderParams{
		CustomerId:   request.GetCustomerId(),
		CustomerName: request.GetCustomerName(),
		OrderItem:    orderItem,
	}
	orderId, err := s.createOrder(ctx, params)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: orderId}, nil
}

func OrderToProto(order *model.Order) *pb.Order {
	protoOrder := &pb.Order{
		Id:           order.ID.Hex(),
		CustomerId:   order.CustomerId,
		CustomerName: order.CustomerName,
		OrderItem:    make([]*pb.OrderItem, len(order.OrderItem)),
	}

	for i, v := range order.OrderItem {
		pbOrderItem := pb.OrderItem{
			ItemId:    int32(v.ItemId),
			Quantity:  int32(v.Quantity),
			UnitPrice: int32(v.UnitPrice),
		}
		protoOrder.OrderItem[i] = &pbOrderItem
	}

	return protoOrder
}
