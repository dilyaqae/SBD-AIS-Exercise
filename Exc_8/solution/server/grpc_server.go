package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions

// prepopulating the drinks
// drinks is a list of pointers to pb.Drink. each drink corresponds to Drink message in .proto
// "in memorz database", server returns this list when client calls GetDrinks
var drinks = []*pb.Drink{
	{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
	{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
	{Id: 3, Name: "Coffee", Price: 2, Description: "Mifare isn't that secure"},
}

// orders storage, nextOrderID keeps track of unique ID for every new order, used by OrderDrink()
var orders []*pb.Order
var nextOrderID int32 = 1

// GetDrinks(), it receives ctx (context) and empty request (as client sends no data)
// it returns DrinksResponse containing drinks list
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinksResponse, error) {
	return &pb.DrinksResponse{Drinks: drinks}, nil
}

// OrderDrink(), it receives ctx (context) and a request containing a list of items the client wants to order
// the server creates a new Order with a unique ID and stores it in-memory
// it returns BoolValue(true) to tell the client the order was successfully stored
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderDrinkRequest) (*wrapperspb.BoolValue, error) {
	// create a new order, assign next available ID and copy all ordered items from client request
	order := &pb.Order{
		Id:    nextOrderID,
		Items: req.Items,
	}
	// increment ID
	nextOrderID++
	// store the order
	orders = append(orders, order)

	// return success
	return wrapperspb.Bool(true), nil
}

// GetOrders(), it receives empty request and ctx (context)
// it returns all orders ever placed
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrdersResponse, error) {
	return &pb.OrdersResponse{Orders: orders}, nil
}
