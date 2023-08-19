package grpc

import (
	"context"
	"eda-in-golang/depot/depotpb/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"eda-in-golang/depot/internal/application"
	"eda-in-golang/depot/internal/application/commands"
)

type server struct {
	app application.App
	depotpbv1.UnimplementedDepotServiceServer
}

var _ depotpbv1.DepotServiceServer = (*server)(nil)

func Register(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	depotpbv1.RegisterDepotServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateShoppingList(ctx context.Context, request *depotpbv1.CreateShoppingListRequest) (*depotpbv1.CreateShoppingListResponse, error) {
	id := uuid.New().String()

	items := make([]commands.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomain(item))
	}

	err := s.app.CreateShoppingList(ctx, commands.CreateShoppingList{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	return &depotpbv1.CreateShoppingListResponse{Id: id}, err
}

func (s server) CancelShoppingList(ctx context.Context, request *depotpbv1.CancelShoppingListRequest) (*depotpbv1.CancelShoppingListResponse, error) {
	err := s.app.CancelShoppingList(ctx, commands.CancelShoppingList{
		ID: request.GetId(),
	})

	return &depotpbv1.CancelShoppingListResponse{}, err
}

func (s server) AssignShoppingList(ctx context.Context, request *depotpbv1.AssignShoppingListRequest) (*depotpbv1.AssignShoppingListResponse, error) {
	err := s.app.AssignShoppingList(ctx, commands.AssignShoppingList{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})
	return &depotpbv1.AssignShoppingListResponse{}, err
}

func (s server) CompleteShoppingList(ctx context.Context, request *depotpbv1.CompleteShoppingListRequest) (*depotpbv1.CompleteShoppingListResponse, error) {
	err := s.app.CompleteShoppingList(ctx, commands.CompleteShoppingList{ID: request.GetId()})
	return &depotpbv1.CompleteShoppingListResponse{}, err
}

func (s server) itemToDomain(item *depotpbv1.OrderItem) commands.OrderItem {
	return commands.OrderItem{
		StoreID:   item.GetStoreId(),
		ProductID: item.GetProductId(),
		Quantity:  int(item.GetQuantity()),
	}
}
