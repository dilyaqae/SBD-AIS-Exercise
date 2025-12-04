package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()

	// 1. List drinks
	fmt.Println("Requesting drinks")
	drinksResp, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%d  description:%q\n",
			d.Id, d.Name, d.Price, d.Description)
	}

	// helper
	findID := func(name string) int32 {
		for _, d := range drinksResp.Drinks {
			if d.Name == name {
				return d.Id
			}
		}
		return 0
	}

	// 2. order drinks
	fmt.Println("Ordering drinks")
	req1 := &pb.OrderDrinkRequest{
		Items: []*pb.OrderItem{
			{DrinkId: findID("Spritzer"), Quantity: 2},
			{DrinkId: findID("Beer"), Quantity: 2},
			{DrinkId: findID("Coffee"), Quantity: 2},
		},
	}
	if _, err := c.client.OrderDrink(ctx, req1); err != nil {
		return err
	}
	fmt.Println("\t> Ordering: 2 x Spritzer")
	fmt.Println("\t> Ordering: 2 x Beer")
	fmt.Println("\t> Ordering: 2 x Coffee")

	// 3. order more
	fmt.Println("Ordering another round of drinks")
	req2 := &pb.OrderDrinkRequest{
		Items: []*pb.OrderItem{
			{DrinkId: findID("Spritzer"), Quantity: 6},
			{DrinkId: findID("Beer"), Quantity: 6},
			{DrinkId: findID("Coffee"), Quantity: 6},
		},
	}
	if _, err := c.client.OrderDrink(ctx, req2); err != nil {
		return err
	}
	fmt.Println("\t> Ordering: 6 x Spritzer")
	fmt.Println("\t> Ordering: 6 x Beer")
	fmt.Println("\t> Ordering: 6 x Coffee")

	// 4. totals
	fmt.Println("Getting the bill")
	ordersResp, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	totals := map[int32]int32{}
	for _, order := range ordersResp.Orders {
		for _, item := range order.Items {
			totals[item.DrinkId] += item.Quantity
		}
	}

	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> Total: %d x %s\n", totals[d.Id], d.Name)
	}

	return nil
}
