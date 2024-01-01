package interfaces

import (
	"cartify/order/interfaces/service"

	pb "cartify/order/proto/order"

	"cartify/order/app/usecase/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"cartify/order/domain/repository"
)

type ServerParams struct {
	OrderRepository repository.OrderRepository
	EventRepository repository.EventRepository
}

func NewServer(params ServerParams) *grpc.Server {
	server := grpc.NewServer()

	orderService := service.NewOrderServer(
		order.NewListOrders(params.OrderRepository),
		order.NewGetOrder(params.OrderRepository),
		order.NewCreateOrder(params.OrderRepository, params.EventRepository),
	)

	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	healthpb.RegisterHealthServer(server, hsrv)

	reflection.Register(server)

	pb.RegisterOrderServiceServer(server, orderService)

	return server

}
